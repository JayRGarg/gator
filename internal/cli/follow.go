package cli

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"context"
	"encoding/json"
	"github.com/jayrgarg/gator/internal/state"
	"github.com/jayrgarg/gator/internal/database"
)

func HandleFollow(s *state.State, cmd Command, user database.User) error {

    if len(cmd.Args) != 1 {
        return fmt.Errorf("Expected 1 argument, got: %v", len(cmd.Args))
    }
	url := cmd.Args[0]

	feed, err := s.Db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting Feed from DB, %v\n", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID: 			uuid.New(),
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
		UserID: 		user.ID,
		FeedID: 		feed.ID,
	}

	feedFollowRow, err := s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Error creating feed follow in DB, %v\n", err)
	}

	jfeed, err := json.MarshalIndent(feedFollowRow, "", "\t")
	if err != nil {
		return fmt.Errorf("Error marshall indenting feed follow row, %v\n", err)
	}
	fmt.Println(string(jfeed))

	return nil
}
