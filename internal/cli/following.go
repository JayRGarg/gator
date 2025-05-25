package cli

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/jayrgarg/gator/internal/state"
	"github.com/jayrgarg/gator/internal/database"
)

func HandleFollowing(s *state.State, cmd Command, user database.User) error {

    if len(cmd.Args) != 0 {
        return fmt.Errorf("Expected 0 arguments, got: %v", len(cmd.Args))
    }

	feedFollowsForUser, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting Feed Follows for User from DB, %v\n", err)
	}
	
	for _, feedFollow := range feedFollowsForUser {
		jfeed, err := json.MarshalIndent(feedFollow, "", "\t")
		if err != nil {
			return fmt.Errorf("Error marshall indenting feed, %v\n", err)
		}
		fmt.Println(string(jfeed))
	}

	return nil
}
