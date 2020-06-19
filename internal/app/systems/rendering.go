package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/gen2brain/raylib-go/raylib"
	"time"
)

type rendering struct {
	background   rl.Color
	camera       rl.Camera2D
	images       map[string]*rl.Image
	textures     map[string]rl.Texture2D
	title        string
	windowHeight int32
	windowWidth  int32
	done         chan bool
}

// NewRendering ...
func NewRendering(width, height int32, title string, background rl.Color) ecs.System {
	return &rendering{
		background:   background,
		images:       map[string]*rl.Image{},
		textures:     map[string]rl.Texture2D{},
		title:        title,
		windowHeight: height,
		windowWidth:  width,
	}
}

// Process ...
func (s *rendering) Process(entityManager *ecs.EntityManager) {
	rl.BeginDrawing()
	rl.BeginMode2D(s.camera)
	rl.ClearBackground(s.background)
	// Render all entities with a position and size.
	for _, e := range entityManager.FilterBy("position", "size") {
		isTexturePresent := s.renderTextureIfPresent(e)
		if !isTexturePresent {
			s.renderBoundingBox(e)
		}
	}
	// Fadeout text if timeout is reached.
	for _, e := range entityManager.FilterBy("text", "timeout") {
		text := e.Get("text").(*components.Text)
		timeout := e.Get("timeout").(*components.Timeout)
		if time.Since(timeout.CreationTime) > timeout.Duration {
			text.Content = ""
		}
	}
	// Ensure that text will always drawn on top.
	for _, e := range entityManager.FilterBy("text") {
		s.renderTextIfPresent(e)
	}
	s.toggleFullscreenIfPresent()
	rl.EndMode2D()
	rl.EndDrawing()
}

// Setup ...
func (s *rendering) Setup() {
	rl.InitWindow(s.windowWidth, s.windowHeight, s.title)
	s.camera = rl.NewCamera2D(
		rl.NewVector2(float32(s.windowWidth/2), float32(s.windowHeight/2)),
		rl.NewVector2(float32(s.windowWidth/2), float32(s.windowHeight/2)),
		0.0,
		1.0,
	)
	rl.SetTargetFPS(60)
}

// Teardown ...
func (s *rendering) Teardown() {
	for _, img := range s.images {
		rl.UnloadImage(img)
	}
	for _, tx := range s.textures {
		rl.UnloadTexture(tx)
	}
	rl.CloseWindow()
}

func (s *rendering) renderBoundingBox(entity *ecs.Entity) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	rl.DrawRectangleLines(
		int32(position.X),
		int32(position.Y),
		int32(size.Width),
		int32(size.Height),
		rl.RayWhite,
	)
}

func (s *rendering) renderTextIfPresent(entity *ecs.Entity) (present bool) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	// Return if text is not present.
	text := entity.Get("text")
	if text == nil {
		return false
	}
	var x, y int32
	txt := text.(*components.Text)
	if !txt.IsEnabled {
		return false
	}
	txtLength := rl.MeasureText(txt.Content, txt.FontSize)
	switch txt.Align {
	case components.TextAlignBottom:
		x = int32(position.X) + int32(size.Width)/2 - txtLength/2
		y = int32(position.Y) + int32(size.Height) + 4
	case components.TextAlignCenter:
		x = int32(position.X) + int32(size.Width)/2 - txtLength/2
		y = int32(position.Y) + int32(size.Height)/2 + 4
	case components.TextAlignLeft:
		x = int32(position.X) - txtLength - 4
		y = int32(position.Y) + int32(size.Height)/2
	case components.TextAlignRight:
		x = int32(position.X) + int32(size.Width) + 4
		y = int32(position.Y) + int32(size.Height)/2
	case components.TextAlignTop:
		x = int32(position.X) + int32(size.Width)/2 - txtLength/2
		y = int32(position.Y) - 4
	}
	rl.DrawText(
		txt.Content,
		x,
		y,
		txt.FontSize,
		txt.Color,
	)
	return true
}

func (s *rendering) renderTextureIfPresent(entity *ecs.Entity) (present bool) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	// Return if texture is not present.
	texture := entity.Get("texture")
	if texture == nil {
		return false
	}
	// Get texture from cache or load it from the filesystem into the cache.
	tex := texture.(*components.Texture)
	fileName := tex.Filename
	if fileName == "" || !tex.IsEnabled {
		return false
	}
	tx, exists := s.textures[fileName]
	if !exists {
		s.textures[fileName] = rl.LoadTexture(fileName)
		tx = s.textures[fileName]
	}
	rl.DrawTextureRec(
		tx,
		rl.NewRectangle(0, 0, size.Width, size.Height),
		rl.NewVector2(position.X, position.Y),
		rl.RayWhite,
	)
	return true
}

func (s *rendering) toggleFullscreenIfPresent() {
	if rl.IsKeyPressed(rl.KeyT) {
		rl.ToggleFullscreen()
	}
}
