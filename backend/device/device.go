package device

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/AndreWongZH/iothome/models"
	"github.com/AndreWongZH/iothome/socket"
	"github.com/AndreWongZH/iothome/wled"
)

func QueryDevStatus(url string) models.DeviceStatus {
	var devStatus models.DeviceStatus
	var wledConfig wled.WledConfig

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := client.Get(url)
	if err != nil {
		devStatus.Connected = false
		devStatus.On_state = false
		return devStatus
	}
	defer resp.Body.Close()

	// for wled
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		devStatus.Connected = false
		devStatus.On_state = false
		return devStatus
	}

	json.Unmarshal(body, &wledConfig)

	devStatus.Connected = true

	if wledConfig.State.On {

		devStatus.On_state = true
	} else {
		devStatus.On_state = false
	}

	return devStatus
}

type WebsocketMessage struct {
	DevStatuses map[string]models.DeviceStatus `json:"devstatuses"`
	RoomName    string                         `json:"roomname"`
}

func QueryRoomDevices(devList []models.RegisteredDevice, roomName string) {
	devStatuses := make(map[string]models.DeviceStatus)

	for _, dev := range devList {
		devStatuses[dev.Ipaddr] = QueryDevStatus("http://" + dev.Ipaddr + "/json")
	}

	// write to database

	socket.BroadcastMsg(WebsocketMessage{
		DevStatuses: devStatuses,
		RoomName:    roomName,
	})
}
