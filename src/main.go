package main

import (
	"github.com/gin-gonic/gin"
	"hvalfangst/gin-api-with-auth/src/common/db"
	"hvalfangst/gin-api-with-auth/src/common/utils"
	"hvalfangst/gin-api-with-auth/src/users"
	"log"
)

func main() {
	r := gin.Default()

	// Map environment variables from 'config.json' to be used as Configuration
	config, err := utils.ReadConfig("src/config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Connect to the database based on Configuration derived from 'config.json'
	database := db.ConnectDatabase(config)
	defer db.CloseDatabase(database)

	// Create the 'users' table
	err = db.CreateTable(database, (*users.User)(nil))
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	users.ConfigureRoute(r, database)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
