package entities

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// NewPlayer creates a new player with an id on a specific position x and y with a custom width and height.
func NewPlayer(id string, x, y, width, height float32) (player *ecs.Entity) {
	return ecs.NewEntity(id,
		[]ecs.Component{
			&components.Position{
				X: x,
				Y: y,
			},
			&components.Input{

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
