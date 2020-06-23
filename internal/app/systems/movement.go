package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// movement ...
type movement struct{}

// NewMovement ...
func NewMovement() ecs.System {
	return &movement{}
}

// Process ...
func (s *movement) Process(entityManager *ecs.EntityManager) (state int) {
	for _, e := range entityManager.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
	return ecs.StateEngineContinue
}

// Setup ...
func (s *movement) Setup() {}

// Teardown ...
func (s *movement) Teardown() {}
