package routine

import (
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/users/repository"
	"log"
	"time"
)

// StartUserDeletionRoutine Define a function to periodically check for users marked for deletion and delete them
func StartUserDeletionRoutine(db *pg.DB) {
	// Create a ticker that ticks every 10 seconds
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Get a list of users marked for deletion
			usersForDeletion, err := repository.GetUsersForDeletion(db)
			if err != nil {
				log.Printf("Error retrieving users marked for deletion: %v", err)
				continue
			}

			// Delete users marked for deletion
			for _, user := range usersForDeletion {
				if err := repository.DeleteByID(db, user.ID); err != nil {
					log.Printf("Error deleting user %s: %v", user.Email, err)
				} else {
					log.Printf("User %s marked for deletion has been deleted.", user.Email)
				}
			}
		}
	}
}
