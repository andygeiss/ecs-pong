package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// AI ...
type AI struct{}

// NewAI ...
func NewAI() ecs.System {
	return &AI{}
}

// Process ...
func (s *AI) Process(entityManager *ecs.EntityManager) {
	for _, e := range entityManager.FilterBy("ai", "position", "velocity") {
		ball := entityManager.Get("ball")
		s.handleBallPosition(e, ball)
	}
}

// Setup ...
func (s *AI) Setup() {}

// Teardown ...
func (s *AI) Teardown() {}

func (s *AI) handleBallPosition(entity, ball *ecs.Entity) {
	ai := entity.Get("ai").(*components.AI)
	ballPosition := ball.Get("position").(*components.Position)
	position := entity.Get("position").(*components.Position)
	velocity := entity.Get("velocity").(*components.Velocity)
	if position.Y+velocity.Y < ballPosition.Y {
		ai.Down = true
		ai.Up = false
	}
	if position.Y+velocity.Y > ballPosition.Y {
		ai.Down = false
		ai.Up = true
	}
	if ai.Down {
		velocity.Y = 4
	}
	if ai.Up {
		velocity.Y = -4
	}
}
