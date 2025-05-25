package cli

import (
	"fmt"
	"context"
	"github.com/jayrgarg/gator/internal/state"
	"github.com/jayrgarg/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(*state.State, Command, database.User) error) (func(*state.State, Command) error) {
	return func(s *state.State, cmd Command) error {
		currentUserName := s.Conf.CurrentUserName
		user, err := s.Db.GetUser(context.Background(), currentUserName)
		if err != nil {
			return fmt.Errorf("Error getting current User from DB, %v\n", err)
		}
		return handler(s, cmd, user)
	}
}
