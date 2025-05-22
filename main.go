package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/jayrgarg/gator/internal/database"
	"github.com/jayrgarg/gator/internal/cli"
	"github.com/jayrgarg/gator/internal/config"
	"github.com/jayrgarg/gator/internal/state"
)

import _ "github.com/lib/pq"

func main() {
    conf, err := config.Read()
    if err != nil {
        fmt.Printf("Failed to get configuration: %v\n", err)
        os.Exit(1)
        return
    }

	db, err := sql.Open("postgres", conf.DbUrl)
	dbQueries := database.New(db)


	st := &state.State{ Db: dbQueries, Conf: &conf }
    cmds := cli.Commands{ CmdMap: make(map[string]func(*state.State, cli.Command) error) }
    cmds.Register("login", cli.HandlerLogin)
    cmds.Register("register", cli.HandlerRegister)
    // for k, _ := range cmds.CmdMap {
    //     fmt.Println(k, "value is")
    // }


    userArgs := os.Args
    if len(userArgs) < 2 {
        fmt.Printf("Expected at least 2 arguments, got %v\n", len(userArgs))
        os.Exit(1)
    }
    cmd := cli.Command {
        Name: userArgs[1],
        Args: userArgs[2:],
    }

    // fmt.Printf("Args: %v, %v\n", userArgs[0], userArgs[1])

    err = cmds.Run(st, cmd)
    if err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }

    // err = conf.SetUser("jayrgarg")
    // if err != nil {
    //     fmt.Printf("Failed to set user: %v\n", err)
    //     return
    // }
    // conf, err = config.Read()
    // if err != nil {
    //     fmt.Printf("Failed to get configuration after setting user: %v\n", err)
    //     return
    // }
    // fmt.Printf("Current Configuration: %+v\n", conf)

    os.Exit(0)
}
