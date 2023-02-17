package models

type RegisteredDevice struct {
	Hostname string `json:"hostname"`
	Ipaddr   string `json:"ipaddr"`
	Nickname string `json:"nickname"`
	Type     string `json:"type"`
}

const (
	Computer string = "computer"
	Switch   string = "switch"
	Wled     string = "wled"
)

type DeviceStatus struct {
	Status bool `json:"status"`
	On     bool `json:"on"`
}
