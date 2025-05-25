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

func HandleAddFeed(s *state.State, cmd Command, user database.User) error {

    if len(cmd.Args) != 2 {
        return fmt.Errorf("Expected 2 arguments, got: %v", len(cmd.Args))
    }
	
	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	feedParams := database.CreateFeedParams{
		ID: 			uuid.New(),
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
		Name: 			feedName,
		Url: 			feedUrl,
		UserID: 		user.ID,
	}

	feed, err := s.Db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("Error creating feed in DB, %v\n", err)
	}
	
	feedFollowParams := database.CreateFeedFollowParams{
		ID: 			uuid.New(),
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
		UserID: 		user.ID,
		FeedID: 		feed.ID,
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Error creating feed follow in DB, %v\n", err)
	}

	jfeed, err := json.MarshalIndent(feed, "", "\t")
	if err != nil {
		return fmt.Errorf("Error marshall indenting feed, %v\n", err)
	}
	fmt.Println(string(jfeed))

	return nil
}

