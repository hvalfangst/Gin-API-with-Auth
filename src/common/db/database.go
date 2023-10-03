package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"hvalfangst/gin-api-with-auth/src/common/utils"
	"log"
)

func ConnectDatabase(config utils.Configuration) *pg.DB {
	opts := &pg.Options{
		User:     config.Db.User,
		Password: config.Db.Password,
		Addr:     config.Db.Address,
		Database: config.Db.Database,
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
