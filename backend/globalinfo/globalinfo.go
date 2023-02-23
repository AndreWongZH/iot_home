package globalinfo

import (
	"database/sql"

	"github.com/AndreWongZH/iothome/models"
)

type ServerState struct {
	Rooms map[string]models.Room
	db    *sql.DB
}

var ServerInfo ServerState

func InitializeGlobals(db *sql.DB) {
	ServerInfo.Rooms = make(map[string]models.Room)
	ServerInfo.db = db
}
