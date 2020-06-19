package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
)

// Movement ...
type Movement struct{}

// NewMovement ...
func NewMovement() ecs.System {
	return &Movement{}
}

// Process ...
func (s *Movement) Process(entityManager *ecs.EntityManager) {
	for _, e := range entityManager.FilterBy("position", "velocity") {
		position := e.Get("position").(*components.Position)
		velocity := e.Get("velocity").(*components.Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
}

// Setup ...
func (s *Movement) Setup() {}

// Teardown ...
func (s *Movement) Teardown() {}
