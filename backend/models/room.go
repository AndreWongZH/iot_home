package models

type Room struct {
	Name       string                      `json:"name"`
	Devices    map[string]RegisteredDevice `json:"devices"`
	DeviceInfo map[string]DeviceStatus     `json:"deviceInfo"`
}

// type RoomResponse struct {
// 	Name string ``
// }
