package raylib

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/components"
	"github.com/gen2brain/raylib-go/raylib"
)

// input ...
type input struct{}

// NewInput ...
func NewInput() ecs.System {
	return &input{}
}

// Process ...
func (s *input) Process(entityManager *ecs.EntityManager) (state int) {
	for _, e := range entityManager.FilterByMask(components.MaskInput | components.MaskVelocity) {
		s.handleInput(e)
	}
	return ecs.StateEngineContinue
}

// Setup ...
func (s *input) Setup() {}

// Teardown ...
func (s *input) Teardown() {}

func (s *input) handleInput(e *ecs.Entity) {
	input := e.Get(components.MaskInput).(*components.Input)
	velocity := e.Get(components.MaskVelocity).(*components.Velocity)
	input.Down = rl.IsKeyDown(rl.KeyS)
	input.Up = rl.IsKeyDown(rl.KeyW)
	velocity.Y = 0
	if input.Down {
		velocity.Y = 4
	}
	if input.Up {
		velocity.Y = -4
	}
}
