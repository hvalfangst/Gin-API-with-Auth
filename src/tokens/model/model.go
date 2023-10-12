package model

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	ID           uuid.UUID `pg:"type:uuid,pk"`
	CreationDate time.Time
	UserID       int64
	Activities   []*TokenActivity
	_            struct{} `pg:"_schema:tokens"`
}

type TokenActivity struct {
	ID       int64
	TokenID  uuid.UUID `pg:"type:uuid"`
	Endpoint string
	UsedAt   time.Time
	_        struct{} `pg:"_schema:token_activities"`
}
