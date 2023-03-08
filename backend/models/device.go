package models

type RegisteredDevice struct {
	Hostname string `json:"hostname"`
	Ipaddr   string `json:"ipaddr"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

const (
	Computer string = "computer"
	Switch   string = "switch"
	Wled     string = "wled"
)

type DeviceStatus struct {
	Connected bool `json:"connected"`
	On_state  bool `json:"on_state"`
}
