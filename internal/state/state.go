package state

import (
	"github.com/jayrgarg/gator/internal/config"
	"github.com/jayrgarg/gator/internal/database"
)

type State struct {
	Db 			*database.Queries
    Conf        *config.Config
}
