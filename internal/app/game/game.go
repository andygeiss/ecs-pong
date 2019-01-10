package game

import (
	"github.com/andygeiss/ecs"
	myEntities "github.com/andygeiss/ecs-pong/internal/app/entities"
	mySystems "github.com/andygeiss/ecs-pong/internal/app/systems"
	"github.com/andygeiss/ecs/systems"
	"github.com/gen2brain/raylib-go/raylib"
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
		myEntities.NewBall("ball", windowWidth/2, windowHeight/2, ballWidth, ballHeight),
		myEntities.NewPlayer("player", 10, windowHeight/2, paddleWidth, paddleHeight),
		myEntities.NewEnemy("enemy", windowWidth-paddleWidth-10, windowHeight/2, paddleWidth, paddleHeight),
		myEntities.NewScoreboard("scoreboard", 0, 0, windowWidth, 0),
	)
	return
}

// NewSystemsManager ...
func NewSystemsManager() (systemManager *ecs.SystemManager) {
	systemManager = ecs.NewSystemManager()
	systemManager.Add(
		mySystems.NewAI(),
		mySystems.NewInput(),
		mySystems.NewCollision(windowWidth, windowHeight),
		mySystems.NewScore(),
		systems.NewAudio(),
		systems.NewMovement(),
		systems.NewRendering(windowWidth, windowHeight, title, rl.Beige),
	)
	return
}
