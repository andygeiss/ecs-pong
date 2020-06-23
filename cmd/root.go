package cmd

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/entities"
	"github.com/andygeiss/ecs-pong/internal/app/systems"
	"github.com/andygeiss/ecs-pong/internal/app/systems/raylib"
	"github.com/spf13/cobra"
	"log"
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

var rootCmd = &cobra.Command{
	Use:  "ecs-pong",
	Long: `A basic demonstration of an Entity Component System.`,
	Run: func(cmd *cobra.Command, args []string) {
		entityManager := ecs.NewEntityManager()
		entityManager.Add(
			entities.NewBall("ball", windowWidth/2, windowHeight/2, ballWidth, ballHeight),
			entities.NewPlayer("player", 10, windowHeight/2, paddleWidth, paddleHeight),
			entities.NewEnemy("enemy", windowWidth-paddleWidth-10, windowHeight/2, paddleWidth, paddleHeight),
			entities.NewScoreboard("scoreboard", 0, 0, windowWidth, 1),
			entities.NewStatus("status", 0, 0, windowWidth/3, 1),
		)
		systemManager := ecs.NewSystemManager()
		systemManager.Add(
			systems.NewAI(),
			// pixel.NewInput(),
			 raylib.NewInput(),
			systems.NewCollision(windowWidth, windowHeight),
			systems.NewScore(5),
			// pixel.NewAudio(),
			 raylib.NewAudio(),
			systems.NewMovement(),
			// pixel.New2DRendering(windowWidth, windowHeight, title),
			 raylib.NewRendering(windowWidth, windowHeight, title),
		)
		engine := ecs.NewEngine(
			entityManager,
			systemManager,
		)
		engine.Setup()
		engine.Run()
		engine.Teardown()
	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
