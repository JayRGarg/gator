package cli

import (
    "fmt"
	"github.com/jayrgarg/gator/internal/state"
)


func HandlerLogin(s *state.State, cmd Command) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("Expected 1 argument, got: %v", len(cmd.Args))
    }

    userName := cmd.Args[0]

    err := s.Conf.SetUser(userName)
    if err != nil {
        return fmt.Errorf("Error setting user: %w", err)
    }

    fmt.Printf("User has been set to: %v\n", userName)
    return nil
}
