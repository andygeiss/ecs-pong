package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// ai ...
type ai struct{}

// NewAI ...
func NewAI() ecs.System {
	return &ai{}
}

// Process ...
func (s *ai) Process(entityManager *ecs.EntityManager) (state int) {
	ball := entityManager.Get("ball")
	for _, e := range entityManager.FilterBy("ai", "position", "velocity") {
		s.handleBallPosition(e, ball)
	}
	return ecs.StateEngineContinue
}

// Setup ...
func (s *ai) Setup() {}

// Teardown ...
func (s *ai) Teardown() {}

func (s *ai) handleBallPosition(entity, ball *ecs.Entity) {
	ai := entity.Get("ai").(*components.AI)
	position := entity.Get("position").(*components.Position)
	velocity := entity.Get("velocity").(*components.Velocity)
	ballPosition := ball.Get("position").(*components.Position)
	if position.Y+velocity.Y < ballPosition.Y {
		ai.Down = true
		ai.Up = false
	}
	if position.Y+velocity.Y > ballPosition.Y {
		ai.Down = false
		ai.Up = true
	}
	if ai.Down {
		velocity.Y = 2
	}
	if ai.Up {
		velocity.Y = -2
	}
}
