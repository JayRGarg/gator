package cli

import (
	"fmt"
	"context"
	"github.com/jayrgarg/gator/internal/state"
)

func HandleUsers(s *state.State, cmd Command) error {

    if len(cmd.Args) != 0 {
        return fmt.Errorf("Expected 0 arguments, got: %v", len(cmd.Args))
    }
	names, err := s.Db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Error Getting Users: ", err)
	}
	for _, name := range names {
		fmt.Printf("* %s", name)
		if name == s.Conf.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	return nil
}
