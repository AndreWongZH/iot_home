package models

type Room struct {
	Name    string             `json:"name"`
	Devices []RegisteredDevice `json:"devices"`
}

// type RoomResponse struct {
// 	Name string ``
// }
