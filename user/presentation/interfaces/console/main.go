package main

import (
	"log"

	"github.com/ecommerce/user/infrastructure/persistence"
	"github.com/ecommerce/user/presentation/interfaces/console/commands"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	container := persistence.NewContainer()

	var rootCmd = &cobra.Command{Use: "app"}

	rootCmd.AddCommand(commands.NewCreateServiceTokenCmd(container))

	rootCmd.Execute()
}
