package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"hvalfangst/gin-api-with-auth/src/common/utils/configuration"
	"log"
)

func ConnectDatabase(config configuration.Db) *pg.DB {
	opts := &pg.Options{
		User:     config.User,
		Password: config.Password,
		Addr:     config.Address,
		Database: config.Database,
	}
	return pg.Connect(opts)
}

func CreateTable(db *pg.DB, model interface{}) error {
	err := db.Model(model).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}

func CloseDatabase(db *pg.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
