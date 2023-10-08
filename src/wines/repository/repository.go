package repository

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/wines/model"
	"log"
)

func Create(db *pg.DB, wine *model.Wine) error {
	_, err := db.Model(wine).Insert()
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func Update(db *pg.DB, wineID int64, wine *model.Wine) error {
	_, err := db.Model(wine).Where("id = ?", wineID).Update()
	if err != nil {
		log.Printf("Error updating wine: %v", err)
		return err
	}
	return nil
}

func GetById(db *pg.DB, wineID int64) (*model.Wine, error) {
	wine := &model.Wine{}
	err := db.Model(wine).Where("id = ?", wineID).Select()
	if err != nil {
		log.Printf("Error retrieving wine by ID: %v", err)
		return nil, err
	}
	return wine, nil
}

func Delete(db *pg.DB, wineID int64) error {
	wine := &model.Wine{ID: wineID}

	_, err := db.Model(wine).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}
