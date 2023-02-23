package globalinfo

// import (
// 	"database/sql"

// 	"github.com/AndreWongZH/iothome/models"
// )

// type ServerState struct {
// 	Rooms map[string]models.Room
// 	db    *sql.DB
// }

// var ServerInfo *ServerState

// func InitializeGlobals(db *sql.DB) {
// 	ServerInfo = &ServerState{
// 		Rooms: make(map[string]models.Room),
// 		db:    db,
// 	}
// 	// ServerInfo.Rooms = make(map[string]models.Room)
// 	// ServerInfo.db = db
// }

// func (g *ServerState) test() {
// 	println("hello world", g.db)
// }
