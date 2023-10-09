package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"hvalfangst/gin-api-with-auth/src/common/security/jwt/tokens/model"
	"log"
)

// - - - - - - - - - - - - - - |TOKENS| - - - - - - - - - - - - - -

func CreateToken(db *pg.DB, token *model.Token) error {
	_, err := db.Model(token).Insert()
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return err
	}
	return nil
}

func GetTokenByID(db *pg.DB, ID uuid.UUID) (*model.Token, error) {
	token := &model.Token{}
	err := db.Model(token).Where("id = ?", ID).Select()
	if err != nil {
		log.Printf("Error retrieving token by ID: %v", err)
		return nil, err
	}
	return token, nil
}

func DeleteToken(db *pg.DB, ID uuid.UUID) error {
	token := &model.Token{ID: ID}
	_, err := db.Model(token).WherePK().Delete()
	if err != nil {
		log.Printf("Error deleting token: %v", err)
		return err
	}
	return nil
}

// - - - - - - - - - - - - - - |TOKEN_USAGES| - - - - - - - - - - - - - -

func CreateTokenUsage(db *pg.DB, tokenUsage *model.TokenUsage) error {
	_, err := db.Model(tokenUsage).Insert()
	if err != nil {
		log.Printf("Error creating TokenUsage entry: %v", err)
		return err
	}
	return nil
}

func GetTokenUsageByID(db *pg.DB, ID uuid.UUID) (*model.TokenUsage, error) {
	token := &model.TokenUsage{}
	err := db.Model(token).Where("id = ?", ID).Select()
	if err != nil {
		log.Printf("Error retrieving token by ID: %v", err)
		return nil, err
	}
	return token, nil
}

func DeleteTokenUsage(db *pg.DB, ID uuid.UUID) error {
	tokenUsage := &model.TokenUsage{}
	_, err := db.Model(tokenUsage).Where("token_id = ?", ID).Delete()
	if err != nil {
		return err
	}
	return nil
}
