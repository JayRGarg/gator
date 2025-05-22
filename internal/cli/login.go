package cli

import (
    "fmt"
	"context"
	"github.com/jayrgarg/gator/internal/state"
)


func HandlerLogin(s *state.State, cmd Command) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("Expected 1 argument (username), got: %v", len(cmd.Args))
    }

    userName := cmd.Args[0]

	user, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
        return fmt.Errorf("Error Getting User from Database: %w", err)
	}

    err = s.Conf.SetUser(user.Name)
    if err != nil {
        return fmt.Errorf("Error setting user: %w", err)
    }

    fmt.Printf("User has been set to: %v\n", user.Name)
    return nil
}
