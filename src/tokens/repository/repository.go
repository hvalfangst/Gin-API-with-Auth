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

func ListTokens(db *pg.DB) ([]*model.Token, error) {
	var tokens []*model.Token
	err := db.Model(&tokens).Select()
	if err != nil {
		log.Printf("Error retrieving tokens: %v", err)
		return nil, err
	}

	// Create a map to store token activities associated with each token
	activitiesMap := make(map[uuid.UUID][]*model.TokenActivity)

	// Fetch all token activities
	var activities []*model.TokenActivity
	err = db.Model(&activities).Select()
	if err != nil {
		log.Printf("Error retrieving token activities: %v", err)
		return nil, err
	}

	// Organize token activities by token ID
	for _, activity := range activities {
		activitiesMap[activity.TokenID] = append(activitiesMap[activity.TokenID], activity)
	}

	// Populate token activities for each token
	for _, token := range tokens {
		token.Activities = activitiesMap[token.ID]
	}

	return tokens, nil
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
	err := db.Model(&tokenActivity).Where("token_id = ?", ID).Select()
	if err != nil {
		log.Printf("Error retrieving token activity for token ID %s: %v", ID.String(), err)
		return nil, err
	}
	return tokenActivity, nil
}

func DeleteTokenActivity(db *pg.DB, ID uuid.UUID) error {
	tokenActivity := &model.TokenActivity{}
	_, err := db.Model(tokenActivity).Where("id = ?", ID).Delete()
	if err != nil {
		return err
	}
	return nil
}
