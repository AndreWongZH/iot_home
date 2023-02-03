package globalinfo

type RegisteredDevice struct {
	Hostname string `json:"hostname"`
	Ipaddr   string `json:"ipaddr"`
	Nickname string `json:"nickname"`
	Type     string `json:"type"`
}

type ServerState struct {
	Devices []RegisteredDevice
}

var ServerInfo ServerState

func InitializeGlobals() {
	ServerInfo.Devices = []RegisteredDevice{}
}

const (
	Computer string = "computer"
	Switch   string = "switch"
	Wled     string = "wled"
)
