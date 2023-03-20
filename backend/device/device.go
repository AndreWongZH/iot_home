package device

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/models"
	"github.com/AndreWongZH/iothome/socket"
	"github.com/AndreWongZH/iothome/wled"
)

func QueryDevStatus(ipAddr string, devType models.DeviceType) models.DeviceStatus {
	var devStatus models.DeviceStatus

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	// for wled
	if devType == models.Wled {
		var wledConfig wled.WledConfig

		resp, err := client.Get("http://" + ipAddr + "/json")
		if err != nil {
			devStatus.Connected = false
			devStatus.On_state = false
			return devStatus
		}
		defer resp.Body.Close()

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

	if devType == models.Switch {
		var tasmotaStatus models.TasmotaStatus

		resp, err := client.Get("http://" + ipAddr + "/cm?cmnd=Status")
		if err != nil {
			devStatus.Connected = false
			devStatus.On_state = false
			return devStatus
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			devStatus.Connected = false
			devStatus.On_state = false
			return devStatus
		}

		json.Unmarshal(body, &tasmotaStatus)
		fmt.Println("succes switch")

		fmt.Println(tasmotaStatus.Status.DeviceName)
		fmt.Println(tasmotaStatus.Status.Power)

		if tasmotaStatus.Status.Power == 1 {
			devStatus.On_state = true
		} else {
			devStatus.On_state = false
		}

		devStatus.Connected = true

		return devStatus
	}

	return devStatus
}

type WebsocketMessage struct {
	DevStatuses map[string]models.DeviceStatus `json:"devstatuses"`
	RoomName    string                         `json:"roomname"`
}

func QueryRoomDevices(devList []models.RegisteredDevice, devStatuses map[string]models.DeviceStatus, roomName string) {
	updatedDevStatuses := make(map[string]models.DeviceStatus)

	for _, dev := range devList {
		updatedDevStatuses[dev.Ipaddr] = QueryDevStatus(dev.Ipaddr, dev.Type)

		if updatedDevStatuses[dev.Ipaddr].Connected != devStatuses[dev.Ipaddr].Connected ||
			updatedDevStatuses[dev.Ipaddr].On_state != devStatuses[dev.Ipaddr].On_state {

			// write to database
			err := database.Dbman.UpdateDevStatus(roomName, dev.Ipaddr, updatedDevStatuses[dev.Ipaddr])
			if err != nil {
				log.Println("failed to update to database")
			}
		}
	}

	socket.BroadcastMsg(WebsocketMessage{
		DevStatuses: updatedDevStatuses,
		RoomName:    roomName,
	})
}

func QueryAllDevices(exit chan bool) {
	for {
		fmt.Println("ping check for all devices")

		roomList, err := database.Dbman.GetRooms()
		if err != nil {
			log.Println("Error getting room list")
		}

		for _, room := range roomList {
			devList, devStatuses, err := database.Dbman.GetDevices(room.Name)
			if err != nil {
				log.Println("Error getting device list")
			}
			QueryRoomDevices(devList, devStatuses, room.Name)
		}

		select {
		case <-exit:
			return
		default:
			time.Sleep(1 * time.Minute)
		}
	}
}
