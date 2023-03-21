package device

import (
	"fmt"
	"log"
	"time"

	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/device/tasmota"
	"github.com/AndreWongZH/iothome/device/wled"
	"github.com/AndreWongZH/iothome/models"
	"github.com/AndreWongZH/iothome/socket"
)

// query the status of a generic device
func QueryDevStatus(ipAddr string, devType models.DeviceType) models.DeviceStatus {
	devStatus := models.DeviceStatus{Connected: false, On_state: false}

	switch devType {
	case models.Wled:
		devStatus = wled.QueryWledStatus(ipAddr)
	case models.Switch:
		devStatus = tasmota.QueryTasmotaStatus(ipAddr)
	}

	return devStatus
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

	socket.BroadcastMsg(socket.DevStatusesMsg{
		DevStatuses: updatedDevStatuses,
		RoomName:    roomName,
	})
}

// a go routine to check the connection of devices every 1 min
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
