package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/models"
	"github.com/AndreWongZH/iothome/wled"
	"github.com/gin-gonic/gin"
)

func sendResultJson(ctx *gin.Context, success bool, err error, errStr string, data interface{}) {
	if success {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    data,
		})
	} else {
		var errorMessage interface{}
		if err != nil {
			errorMessage = err
		} else {
			errorMessage = errStr
		}

		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": success,
			"error":   errorMessage,
		})
	}
}

func getServerStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func createRoom(ctx *gin.Context) {
	var room models.RoomInfo

	err := ctx.BindJSON(&room)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	err = database.Dbman.AddRoom(room)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	sendResultJson(ctx, true, nil, "", nil)
}

func getRooms(ctx *gin.Context) {
	rooms, err := database.Dbman.GetRooms()
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	sendResultJson(ctx, true, nil, "", rooms)
}

func addDevice(ctx *gin.Context) {
	var registeredDevice models.RegisteredDevice

	roomName := ctx.Param("roomname")

	err := ctx.BindJSON(&registeredDevice)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	err = database.Dbman.AddDevice(registeredDevice, roomName)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	sendResultJson(ctx, true, nil, "", nil)
}

func showDevices(ctx *gin.Context) {
	roomName := ctx.Param("roomname")

	devList, devStatus, err := database.Dbman.GetDevices(roomName)

	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"success":   true,
	// 	"devList":   devList,
	// 	"devStatus": devStatus,
	// })

	sendResultJson(ctx, true, nil, "", gin.H{
		"devList":   devList,
		"devStatus": devStatus,
	})
}

func toggleDevice(ctx *gin.Context) {
	roomName := ctx.Param("roomname")
	ipAddr := ctx.Param("ip")
	toggle := ctx.Param("toggle")

	device_id, devInfo, _, err := database.Dbman.GetDevice(roomName, ipAddr)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	if devInfo.Type == "wled" {
		var wledSwitch wled.WledSwitch
		wledSwitch.On = toggle == "on"
		marshalled, err := json.Marshal(wledSwitch)
		if err != nil {
			sendResultJson(ctx, false, err, "", nil)
		}

		resp, err := http.Post("http://"+ipAddr+"/json", "application/json", bytes.NewBuffer(marshalled))
		if err != nil {
			sendResultJson(ctx, false, err, "", nil)
		}

		fmt.Println(resp)
		// need to check if resp is success also

		defer resp.Body.Close()
	}

	on_state := 0
	if toggle == "on" {
		on_state = 1
	}
	err = database.Dbman.UpdateDevStatus(device_id, on_state)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	sendResultJson(ctx, true, nil, "", nil)
}

func getWledConfigs(ctx *gin.Context) {
	var wledConfig wled.WledConfig

	ip := ctx.Param("ip")

	client := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get("http://" + ip + "/json")
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	json.Unmarshal(body, &wledConfig)

	sendResultJson(ctx, true, nil, "", wledConfig.State)
}

func setWled(ctx *gin.Context) {
	// send json to esp device

	ip := ctx.Param("ip")

	var wledState wled.State

	err := ctx.BindJSON(&wledState)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	marshalled, err := json.Marshal(wledState)
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	resp, err := http.Post("http://"+ip+"/json", "application/json", bytes.NewBuffer(marshalled))
	if err != nil {
		sendResultJson(ctx, false, err, "", nil)
	}

	defer resp.Body.Close()

	sendResultJson(ctx, true, nil, "", nil)
}

// func discoverNetworkDevices(ctx *gin.Context) {
// scan for network devices
// using ssdp
// }
