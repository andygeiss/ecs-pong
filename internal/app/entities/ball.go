package entities

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs/components"
)

// NewBall creates a new ball with an id on a specific position x and y with a custom width and height.
func NewBall(id string, x, y, width, height float32) (ball *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Position{X: x, Y: y},
			&components.Size{Width: width, Height: height},
			&components.Sound{
				EventFilename: map[string]string{
					"collision": "assets/sounds/collision.wav",
				},
				Filename: "",
			},
			&components.Texture{Filename: "assets/textures/ball.png"},
			&components.Velocity{X: -3, Y: 2},
		},
	}
}
