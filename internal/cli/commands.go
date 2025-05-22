package cli

import (
    "fmt"
	"github.com/jayrgarg/gator/internal/state"
)

type Command struct {
    Name        string
    Args        []string
}

type Commands struct {
    CmdMap      map[string]func(*state.State, Command) error
}

func (c *Commands) Run(s *state.State, cmd Command) error {
    // fmt.Printf("cmd.Name: %v\n", cmd.Name)
    handler, ok := c.CmdMap[cmd.Name]
    if !ok {
        return fmt.Errorf("Command does not exist")
    } else {
        return handler(s, cmd)
    }
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
    // fmt.Println("name", name)
    c.CmdMap[name] = f
    return
}
