package cli

import (
	"fmt"
	"context"
	"github.com/jayrgarg/gator/internal/state"
	"github.com/jayrgarg/gator/internal/database"
)

func HandleUnfollow(s *state.State, cmd Command, user database.User) error {

    if len(cmd.Args) != 1 {
        return fmt.Errorf("Expected 1 argument, got: %v", len(cmd.Args))
    }
	url := cmd.Args[0]

	feed, err := s.Db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting Feed from DB, %v\n", err)
	}

	params := database.DeleteFeedFollowParams{
		UserID: 		user.ID,
		FeedID: 		feed.ID,
	}

	err = s.Db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error Deleting Feed Follow from DB, %v\n", err)
	}

	return nil
}
