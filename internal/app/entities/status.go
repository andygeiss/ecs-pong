package entities

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	"time"
)

// NewStatus creates a new status message with an id on a specific position x and y with a custom width and height.
func NewStatus(id string, x, y, width, height float32) (player *ecs.Entity) {
	return ecs.NewEntity(id,
		[]ecs.Component{
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
				FontSize:  40,
				IsEnabled: true,
			},
			&components.Timeout{
				CreationTime: time.Now(),
				Duration:     time.Second * 3,
			},
		})
}
