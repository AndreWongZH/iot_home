package models

type Room struct {
	Name       string             `json:"name"`
	Devices    []RegisteredDevice `json:"devices"`
	DeviceInfo map[string]DeviceStatus
}

// type RoomResponse struct {
// 	Name string ``
// }
