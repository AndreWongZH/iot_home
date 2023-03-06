package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/globals"
	"github.com/AndreWongZH/iothome/inputvalid"
	"github.com/AndreWongZH/iothome/models"
	"github.com/AndreWongZH/iothome/wled"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func sendResultJson(ctx *gin.Context, success bool, err error, data interface{}) {
	if success {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    data,
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": success,
			"error":   err,
		})
	}
}

func getServerStatus(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(globals.UserKey)

	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
		"count":  count,
	})
}

func loginPost(ctx *gin.Context) {
	fmt.Println(ctx.GetHeader("Set-Cookie"))
	var userCreds models.UserCreds

	session := sessions.Default(ctx)
	user := session.Get(globals.UserKey)
	if user != nil {
		fmt.Println(user)
		sendResultJson(ctx, false, errors.New("user already logged in"), nil)
		return
	}

	err := ctx.BindJSON(&userCreds)
	if err != nil {
		log.Println(err)
		sendResultJson(ctx, false, err, nil)
		return
	}

	session.Set(globals.UserKey, userCreds.Username)
	if err := session.Save(); err != nil {
		log.Println(err)
		sendResultJson(ctx, false, err, nil)
		return
	}

	user = session.Get(globals.UserKey)
	fmt.Println(user)

	sendResultJson(ctx, true, nil, userCreds)
}

func createRoom(ctx *gin.Context) {
	var room models.RoomInfo

	err := ctx.BindJSON(&room)
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}
	room.Name = strings.Trim(room.Name, " ")

	if err := inputvalid.CheckRoomInput(room.Name); err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	err = database.Dbman.AddRoom(room)
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	sendResultJson(ctx, true, nil, nil)
}

func getRooms(ctx *gin.Context) {
	rooms, err := database.Dbman.GetRooms()
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	sendResultJson(ctx, true, nil, rooms)
}

func addDevice(ctx *gin.Context) {
	var registeredDevice models.RegisteredDevice

	roomName := ctx.Param("roomname")

	err := ctx.BindJSON(&registeredDevice)
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	if err := inputvalid.CheckDeviceInput(&registeredDevice); err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	err = database.Dbman.AddDevice(registeredDevice, roomName)
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	sendResultJson(ctx, true, nil, nil)
}

func showDevices(ctx *gin.Context) {
	roomName := ctx.Param("roomname")

	devList, devStatus, err := database.Dbman.GetDevices(roomName)

	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	sendResultJson(ctx, true, nil, gin.H{
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
		sendResultJson(ctx, false, err, nil)
		return
	}

	if devInfo.Type == "wled" {
		var wledSwitch wled.WledSwitch
		wledSwitch.On = toggle == "on"
		marshalled, err := json.Marshal(wledSwitch)
		if err != nil {
			sendResultJson(ctx, false, err, nil)
			return
		}

		resp, err := http.Post("http://"+ipAddr+"/json", "application/json", bytes.NewBuffer(marshalled))
		if err != nil {
			sendResultJson(ctx, false, err, nil)
			return
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
		sendResultJson(ctx, false, err, nil)
		return
	}

	sendResultJson(ctx, true, nil, nil)
}

func getWledConfigs(ctx *gin.Context) {
	var wledConfig wled.WledConfig

	ip := ctx.Param("ip")

	client := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get("http://" + ip + "/json")
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	json.Unmarshal(body, &wledConfig)

	sendResultJson(ctx, true, nil, wledConfig.State)
}

func setWled(ctx *gin.Context) {
	// send json to esp device

	ip := ctx.Param("ip")

	var wledState wled.State

	err := ctx.BindJSON(&wledState)
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	marshalled, err := json.Marshal(wledState)
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	resp, err := http.Post("http://"+ip+"/json", "application/json", bytes.NewBuffer(marshalled))
	if err != nil {
		sendResultJson(ctx, false, err, nil)
		return
	}

	defer resp.Body.Close()

	sendResultJson(ctx, true, nil, nil)
}

// func discoverNetworkDevices(ctx *gin.Context) {
// scan for network devices
// using ssdp
// }
