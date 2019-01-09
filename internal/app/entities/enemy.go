package entities

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// NewEnemy creates a new player with an id on a specific position x and y with a custom width and height.
func NewEnemy(id string, x, y, width, height float32) (enemy *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.AI{Down: true},
			&components.Position{X: x, Y: y},
			&components.Score{Text: "Enemy wins!", Value: 0, X: 1024/2 - 100, Y: 10},
			&components.Size{Width: width, Height: height},
			&components.Texture{Filename: "assets/textures/paddle.png"},
			&components.Velocity{Y: 0},
		},
	}
}
