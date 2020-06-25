package entities

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// NewBall creates a new ball with an id on a specific position x and y with a custom width and height.
func NewBall(id string, x, y, width, height float32) (ball *ecs.Entity) {
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
			&components.Sound{
				EventFilename: map[string]string{
					"collision": "assets/sounds/collision.wav",
				},
				Filename:  "",
				IsEnabled: true,
				Volume:    1.0,
			},
			&components.Texture{
				Filename:  "assets/textures/ball.png",
				IsEnabled: true,
			},
			&components.Velocity{
				X: -3,
				Y: 2,
			},
		})
}
