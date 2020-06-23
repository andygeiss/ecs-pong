package systems

import (
	"fmt"
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
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
	score := scoreboard.Get(components.MaskScore).(*components.Score)
	scoreboardText := scoreboard.Get(components.MaskText).(*components.Text)
	status := entityManager.Get("status")
	timeout := status.Get(components.MaskTimeout).(*components.Timeout)
	statusText := status.Get(components.MaskText).(*components.Text)
	if score.Enemy >= s.winScore {
		statusText.Content = "Enemy Wins!"
		statusText.Color = 0xff0000ff // red
		timeout.CreationTime = time.Now()
		score.Enemy = 0
		score.Player = 0
	} else if score.Player >= s.winScore {
		statusText.Content = "Player Wins!"
		statusText.Color = 0x00ff00ff // green
		timeout.CreationTime = time.Now()
		score.Enemy = 0
		score.Player = 0
	} else {
		scoreboardText.Content = fmt.Sprintf("%d : %d", score.Player, score.Enemy)
		scoreboardText.Color = 0xffffffff // white
	}
	return ecs.StateEngineContinue
}

// Setup ...
func (s *score) Setup() {}

// Teardown ...
func (s *score) Teardown() {}
