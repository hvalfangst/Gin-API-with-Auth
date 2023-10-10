package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/common/db"
	"hvalfangst/gin-api-with-auth/src/common/utils/configuration"
	tokens "hvalfangst/gin-api-with-auth/src/tokens/model"
	tokensRoute "hvalfangst/gin-api-with-auth/src/tokens/route"
	users "hvalfangst/gin-api-with-auth/src/users/model"
	usersRoute "hvalfangst/gin-api-with-auth/src/users/route"
	wines "hvalfangst/gin-api-with-auth/src/wines/model"
	winesRoute "hvalfangst/gin-api-with-auth/src/wines/route"
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

	// Create the following tables: 'users', 'wines', 'tokens' and 'token_activities'
	createDBTables(err, database)

	// Serve context resources under routes '/users', '/wines', '/tokens' and '/token-activities'
	usersRoute.ConfigureRoute(r, database)
	winesRoute.ConfigureRoute(r, database)
	tokensRoute.ConfigureRoute(r, database)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func createDBTables(err error, database *pg.DB) {

	// Create the 'users' table
	err = db.CreateTable(database, (*users.User)(nil))
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	// Create the 'wines' table
	err = db.CreateTable(database, (*wines.Wine)(nil))
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	// Create the 'tokens' table
	err = db.CreateTable(database, (*tokens.Token)(nil))
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	// Create the 'token_activities' table
	err = db.CreateTable(database, (*tokens.TokenActivity)(nil))
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}
}
