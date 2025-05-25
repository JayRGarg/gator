package cli

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/jayrgarg/gator/internal/state"
)

func HandleFeeds(s *state.State, cmd Command) error {

    if len(cmd.Args) != 0 {
        return fmt.Errorf("Expected 0 arguments, got: %v", len(cmd.Args))
    }
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("Error Getting Feeds: ", err)
	}
	for _, feed := range feeds {
		jfeed, err := json.MarshalIndent(feed, "", "\t")
		if err != nil {
			return fmt.Errorf("Error marshall indenting feed, %v\n", err)
		}
		fmt.Println(string(jfeed))
	}
	return nil
}
