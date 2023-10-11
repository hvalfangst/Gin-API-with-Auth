package model

import "time"

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email" pg:",unique"`
	Password     string    `json:"password"`
	Access       string    `json:"access"`
	Deactivated  bool      `json:"deactivated"`
	ForDeletion  bool      `json:"for_deletion"`
	DeletionTime time.Time `json:"deletion_time"`
	_            struct{}  `pg:"_schema:users"`
}
