package entities

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

// NewStatus creates a new status message with an id on a specific position x and y with a custom width and height.
func NewStatus(id string, x, y, width, height float32) (player *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Position{
				X: x,
				Y: y,
			},
			&components.Size{
				Width:  width,
				Height: height,
			},
			&components.Text{
				Align:     components.TextAlignCenter,
				Color:     rl.Beige,
				FontSize:  40,
				IsEnabled: true,
			},
			&components.Timeout{
				CreationTime: time.Now(),
				Duration:     time.Second * 3,
			},
		},
	}
}
