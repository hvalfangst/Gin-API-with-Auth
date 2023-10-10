package model

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	ID           uuid.UUID `pg:"type:uuid,pk"`
	CreationDate time.Time
	UserID       int64    // Foreign key to User
	_            struct{} `pg:"_schema:tokens"`
}

type TokenActivity struct {
	ID       uuid.UUID `pg:"type:uuid,pk"`
	Endpoint string
	UsedAt   time.Time
	_        struct{} `pg:"_schema:token_activities"`
}
