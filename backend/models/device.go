package models

type DeviceType string

const (
	Wled   DeviceType = "wled"
	Switch DeviceType = "switch"
)

type RegisteredDevice struct {
	Hostname string     `json:"hostname"`
	Ipaddr   string     `json:"ipaddr"`
	Name     string     `json:"name"`
	Type     DeviceType `json:"type"`
}

type DeviceStatus struct {
	Connected bool `json:"connected"`
	On_state  bool `json:"on_state"`
}
