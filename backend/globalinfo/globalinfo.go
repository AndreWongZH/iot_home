package globalinfo

import "github.com/AndreWongZH/iothome/models"

type ServerState struct {
	Rooms map[string]models.Room
}

var ServerInfo ServerState

func InitializeGlobals() {
	ServerInfo.Rooms = make(map[string]models.Room)
}
