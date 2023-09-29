package main

import (
	"fmt"
	"os"
	"rest/config"
	"rest/system"
)

func RunMigrationCLI() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run migrate.go migration:down <filename>")
		return
	}

	// Extract the command and migration filename from the command-line arguments
	command := os.Args[1]
	filename := os.Args[2]

	// Connect to the database
	db := config.InitDb()

	if command == "migration:down" {
		if err := system.RollbackMigration(db, filename); err != nil {
			fmt.Printf("Error rolling back migration '%s': %v\n", filename, err)
		} else {
			fmt.Printf("Migration '%s' has been rolled back\n", filename)
		}
	} else {
		fmt.Println("Invalid command. Use 'migration:down'.")
	}
}
