package entities

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// NewEnemy creates a new player with an id on a specific position x and y with a custom width and height.
func NewEnemy(id string, x, y, width, height float32) (enemy *ecs.Entity) {
	return ecs.NewEntity(
		id,
		[]ecs.Component{
			&components.AI{
				Down: true,
			},
			&components.Position{
				X: x, Y: y,
			},
			&components.Size{
				Width:  width,
				Height: height,
			},
			&components.Texture{
				Filename:  "assets/textures/paddle.png",
				IsEnabled: true,
			},
			&components.Velocity{
				Y: 0,
			},
		})
}
