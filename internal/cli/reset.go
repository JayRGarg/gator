package cli

import (
	"fmt"
	"context"
	"github.com/jayrgarg/gator/internal/state"
)

func HandleReset(s *state.State, cmd Command) error {

    if len(cmd.Args) != 0 {
        return fmt.Errorf("Expected 0 arguments, got: %v", len(cmd.Args))
    }
	err := s.Db.DeleteAllUsers(context.Background())
	return err
}
