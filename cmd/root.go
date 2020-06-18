package cmd

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs-pong/internal/app/game"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:  "ecs-pong",
	Long: `A basic demonstration of an Entity Component System.`,
	Run: func(cmd *cobra.Command, args []string) {
		engine := ecs.NewEngine(
			game.NewEntityManager(),
			game.NewSystemsManager(),
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
