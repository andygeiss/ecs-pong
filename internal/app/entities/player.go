package entities

import (
	"github.com/andygeiss/ecs"
	myComponents "github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/andygeiss/ecs/components"
)

// NewPlayer creates a new player with an id on a specific position x and y with a custom width and height.
func NewPlayer(id string, x, y, width, height float32) (player *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Position{X: x, Y: y},
			&components.Size{Width: width, Height: height},
			&components.Texture{Filename: "assets/textures/paddle.png"},
			&components.Velocity{Y: 0},
			&myComponents.Input{},
		},
	}
}
