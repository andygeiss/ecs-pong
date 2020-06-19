package systems

import (
	"fmt"
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/gen2brain/raylib-go/raylib"
	"time"
)

// score ...
type score struct {
	winScore int32
}

// NewScore ...
func NewScore(winScore int32) ecs.System {
	return &score{
		winScore: winScore,
	}
}

// Process ...
func (s *score) Process(entityManager *ecs.EntityManager) (state int) {
	scoreboard := entityManager.Get("scoreboard")
	score := scoreboard.Get("score").(*components.Score)
	scoreboardText := scoreboard.Get("text").(*components.Text)
	status := entityManager.Get("status")
	timeout := status.Get("timeout").(*components.Timeout)
	statusText := status.Get("text").(*components.Text)
	if score.Enemy >= s.winScore {
		statusText.Content = "Enemy Wins!"
		statusText.Color = rl.Red
		timeout.CreationTime = time.Now()
		score.Enemy = 0
		score.Player = 0
	} else if score.Player >= s.winScore {
		statusText.Content = "Player Wins!"
		statusText.Color = rl.Green
		timeout.CreationTime = time.Now()
		score.Enemy = 0
		score.Player = 0
	} else {
		scoreboardText.Content = fmt.Sprintf("%d : %d", score.Player, score.Enemy)
		scoreboardText.Color = rl.White
	}
	return ecs.StateEngineContinue
}

// Setup ...
func (s *score) Setup() {}

// Teardown ...
func (s *score) Teardown() {}
