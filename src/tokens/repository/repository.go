package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"hvalfangst/gin-api-with-auth/src/tokens/model"
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

func GetToken(db *pg.DB, ID uuid.UUID) (*model.Token, error) {
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

// - - - - - - - - - - - - - - |TOKEN_ACTIVITIES| - - - - - - - - - - - - - -

func CreateTokenUsage(db *pg.DB, tokenUsage *model.TokenActivity) error {
	_, err := db.Model(tokenUsage).Insert()
	if err != nil {
		log.Printf("Error creating TokenActivity entry: %v", err)
		return err
	}
	return nil
}

func GetTokenActivity(db *pg.DB, ID uuid.UUID) ([]*model.TokenActivity, error) {
	var tokenActivity []*model.TokenActivity
	err := db.Model(&tokenActivity).Where("id = ?", ID).Select()
	if err != nil {
		log.Printf("Error retrieving token activity for token ID %s: %v", ID.String(), err)
		return nil, err
	}
	return tokenActivity, nil
}

func DeleteTokenActivity(db *pg.DB, ID uuid.UUID) error {
	tokenActivity := &model.TokenActivity{}
	_, err := db.Model(tokenActivity).Where("token_id = ?", ID).Delete()
	if err != nil {
		return err
	}
	return nil
}
