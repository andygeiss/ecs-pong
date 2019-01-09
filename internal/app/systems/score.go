package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/gen2brain/raylib-go/raylib"
	"log"
	"time"
)

// Score ...
type Score struct{}

// NewScore ...
func NewScore() ecs.System {
	return &Score{}
}

// Process ...
func (s *Score) Process(entityManager *ecs.EntityManager) {
	for _, e := range entityManager.FilterBy("score") {
		if s.entityWins(e) {
			s.showWinnerText(e)
			s.waitAndResetScores(entityManager)
		}
	}
}

// Setup ...
func (s *Score) Setup() {}

// Teardown ...
func (s *Score) Teardown() {}

func (s *Score) entityWins(entity *ecs.Entity) (win bool) {
	score := entity.Get("score").(*components.Score)
	if score.Value >= 10 {
		return true
	}
	return false
}

func (s *Score) showWinnerText(entity *ecs.Entity) {
	score := entity.Get("score").(*components.Score)
	log.Printf("%v", score)
	textSize := int32(50)
	rl.BeginDrawing()
	rl.DrawText(score.Text, score.X, score.Y, textSize, rl.Red)
	rl.EndDrawing()
}

func (s *Score) waitAndResetScores(entityManager *ecs.EntityManager) {
	time.Sleep(time.Second * 3)
	for _, e := range entityManager.FilterBy("score") {
		score := e.Get("score").(*components.Score)
		score.Value = 0
	}
}
