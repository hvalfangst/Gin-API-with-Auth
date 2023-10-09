package model

import (
	"github.com/google/uuid"
	users "hvalfangst/gin-api-with-auth/src/users/model"
	"time"
)

type Token struct {
	ID           uuid.UUID `pg:"type:uuid,pk"`
	CreationDate time.Time
	UserID       int64       // Foreign key to User
	User         *users.User `pg:"rel:has-one"`
}

type TokenUsage struct {
	ID       int64 `json:"id"`
	TokenID  uuid.UUID
	Token    *Token `pg:"rel:has-one"`
	Endpoint string
	UsedAt   time.Time
}
