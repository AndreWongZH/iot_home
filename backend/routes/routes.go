package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AndreWongZH/iothome/globalinfo"
	"github.com/AndreWongZH/iothome/models"
	"github.com/AndreWongZH/iothome/wled"
	"github.com/gin-gonic/gin"
)

func getServerStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func createRoom(ctx *gin.Context) {
	var room models.Room

	err := ctx.BindJSON(&room)
	if err != nil {
		log.Println("failed to bind to json")
		log.Println(err)
		return
	}

	room.Devices = []models.RegisteredDevice{}
	room.DeviceInfo = make(map[string]models.DeviceStatus)
	globalinfo.ServerInfo.Rooms[room.Name] = room
}

func getRooms(ctx *gin.Context) {
	var roomlist = []models.Room{}

	for _, room := range globalinfo.ServerInfo.Rooms {
		roomlist = append(roomlist, room)
	}

	ctx.JSON(http.StatusOK, roomlist)
}

func addDevice(ctx *gin.Context) {
	var registeredDevice models.RegisteredDevice

	roomName := ctx.Param("roomname")

	err := ctx.BindJSON(&registeredDevice)
	if err != nil {
		log.Println("error binding json data to variable")
	}

	if room, ok := globalinfo.ServerInfo.Rooms[roomName]; ok {
		room.Devices = append(room.Devices, registeredDevice)
		room.DeviceInfo[registeredDevice.Ipaddr] = models.DeviceStatus{Status: false, On: false}
		globalinfo.ServerInfo.Rooms[roomName] = room

		fmt.Println("device :", registeredDevice, "is added")
	}
}

func showDevices(ctx *gin.Context) {
	roomName := ctx.Param("roomname")

	ctx.JSON(http.StatusOK, globalinfo.ServerInfo.Rooms[roomName])
}

func getWledConfigs(ctx *gin.Context) {
	var wledConfig wled.WledConfig

	ip := ctx.Param("ip")

	fmt.Println("ip addr to find config is:", ip)

	resp, err := http.Get("http://" + ip + "/json")
	if err != nil {
		log.Println("error retrieving wled configs")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading wled configs")
	}

	json.Unmarshal(body, &wledConfig)

	fmt.Println(wledConfig)
	ctx.JSON(http.StatusOK, wledConfig)
}

func setWled(ctx *gin.Context) {
	// send json to esp device

	ip := ctx.Param("ip")

	var wledState wled.State

	err := ctx.BindJSON(&wledState)
	if err != nil {
		log.Println("error binding to wled state")
	}

	marshalled, err := json.Marshal(wledState)
	if err != nil {
		log.Println("error marshalling data")
	}
	fmt.Println(wledState)

	http.Post("http://"+ip+"/json", "application/json", bytes.NewBuffer(marshalled))

}

func discoverNetworkDevices(ctx *gin.Context) {
	// scan for network devices
	// using ssdp
}
