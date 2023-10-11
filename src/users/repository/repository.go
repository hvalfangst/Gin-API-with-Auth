package repository

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/users/model"
	"log"
)

func Create(db *pg.DB, user *model.User) error {
	_, err := db.Model(user).Insert()
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func GetById(db *pg.DB, userID int64) (*model.User, error) {
	user := &model.User{}
	err := db.Model(user).Where("id = ?", userID).Select()
	if err != nil {
		log.Printf("Error retrieving user by ID: %v", err)
		return nil, err
	}
	return user, nil
}

func GetByEmail(db *pg.DB, email string) (*model.User, error) {
	user := &model.User{}
	err := db.Model(user).Where("email = ?", email).Select()
	if err != nil {
		log.Printf("Error retrieving user by email: %v", err)
		return nil, err
	}
	return user, nil
}

func Update(db *pg.DB, ID int64, user *model.User) error {
	_, err := db.Model(user).Where("id = ?", ID).Update()
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

func DeleteByID(db *pg.DB, userID int64) error {
	user := &model.User{ID: userID}

	_, err := db.Model(user).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}

func DeleteByEmail(db *pg.DB, email string) error {
	user := &model.User{}
	_, err := db.Model(user).Where("email = ?", email).Delete()
	if err != nil {
		return err
	}
	return nil
}
