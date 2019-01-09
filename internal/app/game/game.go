package game

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/entities"
	"github.com/andygeiss/ecs-pong/internal/app/systems"
)

const (
	ballHeight   = 32
	ballWidth    = 32
	paddleHeight = 128
	paddleWidth  = 32
	title        = "ECS Pong"
	windowHeight = 576
	windowWidth  = 1024
)

// NewEntityManager ...
func NewEntityManager() (entityManager *ecs.EntityManager) {
	entityManager = ecs.NewEntityManager()
	entityManager.Add(
		entities.NewBall("ball", windowWidth/2, windowHeight/2, ballWidth, ballHeight),
		entities.NewPlayer("player", 10, windowHeight/2, paddleWidth, paddleHeight),
		entities.NewEnemy("enemy", windowWidth-paddleWidth-10, windowHeight/2, paddleWidth, paddleHeight),
	)
	return
}

// NewSystemsManager ...
func NewSystemsManager() (systemManager *ecs.SystemManager) {
	systemManager = ecs.NewSystemManager()
	systemManager.Add(
		systems.NewAI(),
		systems.NewInput(),
		systems.NewCollision(windowWidth, windowHeight),
		systems.NewMovement(),
		systems.NewRendering(windowWidth, windowHeight, title),
		systems.NewScore(),
		systems.NewAudio(),
	)
	return
}
