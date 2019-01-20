package systems

import (
	"fmt"
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/gen2brain/raylib-go/raylib"
)

// Score ...
type Score struct{}

// NewScore ...
func NewScore() ecs.System {
	return &Score{}
}

// Process ...
func (s *Score) Process(entityManager *ecs.EntityManager) {
	if rl.WindowShouldClose() {
		return
	}
	if ecs.ShouldEnginePause {
		return
	}
	scoreboard := entityManager.Get("scoreboard")
	score := scoreboard.Get("score").(*components.Score)
	text := scoreboard.Get("text").(*components.Text)
	if score.Enemy >= 10 {
		text.Content = "Enemy Wins!"
		text.Color = rl.Red
		score.Enemy = 0
		score.Player = 0
		ecs.ShouldEnginePause = true
	} else if score.Player >= 10 {
		text.Content = "Player Wins!"
		text.Color = rl.Green
		score.Enemy = 0
		score.Player = 0
		ecs.ShouldEnginePause = true
	} else {
		text.Content = fmt.Sprintf("%d : %d", score.Player, score.Enemy)
		text.Color = rl.White
	}
}

// Setup ...
func (s *Score) Setup() {}

// Teardown ...
func (s *Score) Teardown() {}
