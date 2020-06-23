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
	for _, e := range entityManager.FilterByMask(components.MaskAI | components.MaskPosition | components.MaskVelocity) {
		s.handleBallPosition(e, ball)
	}
	return ecs.StateEngineContinue
}

// Setup ...
func (s *ai) Setup() {}

// Teardown ...
func (s *ai) Teardown() {}

func (s *ai) handleBallPosition(entity, ball *ecs.Entity) {
	ai := entity.Get(components.MaskAI).(*components.AI)
	position := entity.Get(components.MaskPosition).(*components.Position)
	velocity := entity.Get(components.MaskVelocity).(*components.Velocity)
	ballPosition := ball.Get(components.MaskPosition).(*components.Position)
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
