package systems

import (
	"fmt"
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/gen2brain/raylib-go/raylib"
)

type rendering struct {
	images       map[string]*rl.Image
	textures     map[string]rl.Texture2D
	title        string
	windowHeight int32
	windowWidth  int32
}

// NewRendering ...
func NewRendering(width, height int32, title string) ecs.System {
	return &rendering{
		images:       map[string]*rl.Image{},
		textures:     map[string]rl.Texture2D{},
		title:        title,
		windowHeight: height,
		windowWidth:  width,
	}
}

// Process ...
func (s *rendering) Process(entityManages *ecs.EntityManager) {
	if rl.WindowShouldClose() {
		ecs.ShouldEngineStop = true
		return
	}
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	for _, e := range entityManages.FilterBy("position", "size") {
		s.renderScoreIfPresent(e)
		present := s.renderTextureIfPresent(e, s.textures)
		if !present {
			s.renderBoundingBox(e)
		}
	}
	rl.EndDrawing()
}

// Setup ...
func (s *rendering) Setup() {
	rl.InitWindow(s.windowWidth, s.windowHeight, s.title)
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

func (s *rendering) renderBoundingBox(entity *ecs.Entity) (present bool) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	rl.DrawRectangleLines(
		int32(position.X),
		int32(position.Y),
		int32(size.Width),
		int32(size.Height),
		rl.RayWhite,
	)
	return true
}

func (s *rendering) renderScoreIfPresent(entity *ecs.Entity) {
	position := entity.Get("position").(*components.Position)
	score := entity.Get("score")
	if score == nil {
		return
	}
	var fontSize int32 = 40
	if position.X < 100 {
		rl.DrawText(
			fmt.Sprintf("%d", score.(*components.Score).Value),
			100, 10, fontSize, rl.White,
		)
		return
	}
	text := fmt.Sprintf("%d", score.(*components.Score).Value)
	textSize := rl.MeasureText(text, fontSize)
	rl.DrawText(
		text,
		s.windowWidth-textSize-100, 10, fontSize, rl.White,
	)
}

func (s *rendering) renderTextureIfPresent(entity *ecs.Entity, textures map[string]rl.Texture2D) (present bool) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	texture := entity.Get("texture")
	if texture == nil {
		return false
	}
	fileName := texture.(*components.Texture).Filename
	tx, exists := textures[fileName]
	if !exists {
		textures[fileName] = rl.LoadTexture(fileName)
	}
	rl.DrawTextureRec(
		tx,
		rl.NewRectangle(0, 0, size.Width, size.Height),
		rl.NewVector2(position.X, position.Y),
		rl.RayWhite,
	)
	return true
}
