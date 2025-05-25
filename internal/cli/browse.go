package cli

import (
	"fmt"
	"context"
	"strconv"
	"encoding/json"
	"github.com/jayrgarg/gator/internal/state"
	"github.com/jayrgarg/gator/internal/database"
)

func HandleBrowse(s *state.State, cmd Command, user database.User) error {

    if len(cmd.Args) > 1 {
        return fmt.Errorf("Expected 1 argument, got: %v", len(cmd.Args))
    }

	limit := 2
	var err error
	if len(cmd.Args) == 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Printf("Error converting argument to int, using default limit of 2, %v\n", err)
		}
	}

	params := database.GetPostsForUserParams {
		UserID: 	user.ID,
		Limit: 		int32(limit),
	}
	posts, err := s.Db.GetPostsForUser(context.Background(), params)

	if err != nil {
		fmt.Println("Error Getting User's Posts: ", err)
	}
	for _, post := range posts {
		jpost, err := json.MarshalIndent(post, "", "\t")
		if err != nil {
			return fmt.Errorf("Error marshall indenting post, %v\n", err)
		}
		fmt.Println(string(jpost))
	}

	return nil
}

