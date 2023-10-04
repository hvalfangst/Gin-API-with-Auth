package users

import (
	"github.com/go-pg/pg/v10"
	"log"
)

func create(db *pg.DB, user *User) error {
	_, err := db.Model(user).Insert()
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func getByID(db *pg.DB, userID int64) (*User, error) {
	user := &User{}
	err := db.Model(user).Where("id = ?", userID).Select()
	if err != nil {
		log.Printf("Error retrieving user by ID: %v", err)
		return nil, err
	}
	return user, nil
}

func getByEmail(db *pg.DB, email string) (*User, error) {
	user := &User{}
	err := db.Model(user).Where("email = ?", email).Select()
	if err != nil {
		log.Printf("Error retrieving user by email: %v", err)
		return nil, err
	}
	return user, nil
}

func deleteByID(db *pg.DB, userID int64) error {
	user := &User{ID: userID}

	_, err := db.Model(user).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}

func deleteByEmail(db *pg.DB, email string) error {
	user := &User{}
	_, err := db.Model(user).Where("email = ?", email).Delete()
	if err != nil {
		return err
	}
	return nil
}
