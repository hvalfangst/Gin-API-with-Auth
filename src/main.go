package main

import (
	"github.com/gin-gonic/gin"
	"hvalfangst/gin-api-with-auth/src/common/db"
	"hvalfangst/gin-api-with-auth/src/common/utils/configuration"
	"hvalfangst/gin-api-with-auth/src/users/model"
	"hvalfangst/gin-api-with-auth/src/users/route"
	"log"
)

func main() {
	r := gin.Default()

	// Fetch JSON based on key 'db' for file 'configuration.json' to be used as Db
	conf, err := configuration.Get("db")
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	// Connect to the database based on Configuration derived from 'configuration.json'
	database := db.ConnectDatabase(conf.(configuration.Db))
	defer db.CloseDatabase(database)

	// Create the 'users' table
	err = db.CreateTable(database, (*model.User)(nil))
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	route.ConfigureRoute(r, database)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
