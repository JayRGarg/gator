package cli

import (
    "fmt"
	"os"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/jayrgarg/gator/internal/state"
	"github.com/jayrgarg/gator/internal/database"
)


func HandlerRegister(s *state.State, cmd Command) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("Expected 1 argument (username), got: %v", len(cmd.Args))
    }

    userName := cmd.Args[0]

	userParams := database.CreateUserParams{
		ID: 			uuid.New(),
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
		Name: 			userName,
	}

	user, err := s.Db.CreateUser(context.Background(), userParams)
	if err != nil {
		fmt.Println("Error Creating User: ", err)
		os.Exit(1)
	}

    err = s.Conf.SetUser(user.Name)
    if err != nil {
        return fmt.Errorf("Error Setting User: %w", err)
    }

	fmt.Printf("User has been created!\n")
	fmt.Printf("Id: %v\n", user.ID)
	fmt.Printf("CreatedAt: %v\n", user.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", user.UpdatedAt)
	fmt.Printf("Name: %v\n", user.Name)
    return nil
}

