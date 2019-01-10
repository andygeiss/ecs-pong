package entities

import (
	"github.com/andygeiss/ecs"
	myComponents "github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/andygeiss/ecs/components"
	"github.com/gen2brain/raylib-go/raylib"
)

// NewScoreboard creates a new player with an id on a specific position x and y with a custom width and height.
func NewScoreboard(id string, x, y, width, height float32) (player *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Position{X: x, Y: y},
			&components.Size{Width: width, Height: height},
			&components.Text{Align: components.TextAlignCenter, Color: rl.Beige, FontSize: 40},
			&myComponents.Score{Enemy: 0, Player: 0},
		},
	}
}
