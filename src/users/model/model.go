package model

type User struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email" pg:",unique"`
	Password string   `json:"password"`
	Access   string   `json:"access"`
	_        struct{} `pg:"_schema:users"`
}
