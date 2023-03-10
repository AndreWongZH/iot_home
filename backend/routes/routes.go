package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/device"
	"github.com/AndreWongZH/iothome/globals"
	"github.com/AndreWongZH/iothome/inputvalid"
	"github.com/AndreWongZH/iothome/models"
	"github.com/AndreWongZH/iothome/wled"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func sendResultJson(ctx *gin.Context, success bool, err error, data interface{}, httpCode int) {
	if success {
		ctx.JSON(httpCode, gin.H{
			"success": true,
			"data":    data,
		})

		return
	}

	ctx.AbortWithStatusJSON(httpCode, gin.H{
		"success": success,
		"error":   err.Error(),
	})
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
	var userCreds models.UserCreds

	session := sessions.Default(ctx)
	user := session.Get(globals.UserKey)
	if user != nil {
		fmt.Println(user)
		sendResultJson(ctx, false, errors.New("user already logged in"), nil, http.StatusBadRequest)
		return
	}

	err := ctx.BindJSON(&userCreds)
	if err != nil {
		log.Println(err)
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	session.Set(globals.UserKey, userCreds.Username)
	if err := session.Save(); err != nil {
		log.Println(err)
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	user = session.Get(globals.UserKey)
	fmt.Println(user)

	sendResultJson(ctx, true, nil, nil, http.StatusOK)
}

func logoutPost(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(globals.UserKey)
	log.Println("logging out user:", user)
	if user == nil {
		sendResultJson(ctx, false, errors.New("user token is invalid"), nil, http.StatusBadRequest)
		return
	}

	session.Delete(globals.UserKey)
	if err := session.Save(); err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	sendResultJson(ctx, true, nil, nil, http.StatusOK)
}

func getUsername(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(globals.UserKey)

	if user == nil {
		sendResultJson(ctx, false, errors.New("user token is invalid"), nil, http.StatusBadRequest)
		return
	}

	sendResultJson(ctx, true, nil, user, http.StatusOK)
}

func createRoom(ctx *gin.Context) {
	var room models.RoomInfo

	err := ctx.BindJSON(&room)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}
	room.Name = strings.Trim(room.Name, " ")

	if err := inputvalid.CheckRoomInput(room.Name); err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusBadRequest)
		return
	}

	err = database.Dbman.AddRoom(room)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	sendResultJson(ctx, true, nil, nil, http.StatusOK)
}

func getRooms(ctx *gin.Context) {
	rooms, err := database.Dbman.GetRooms()
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	sendResultJson(ctx, true, nil, rooms, http.StatusOK)
}

func addDevice(ctx *gin.Context) {
	var registeredDevice models.RegisteredDevice

	roomName := ctx.Param("roomname")

	err := ctx.BindJSON(&registeredDevice)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	b, err := database.Dbman.CheckIpExistInRoom(registeredDevice.Ipaddr, roomName)
	if err != nil || b {
		sendResultJson(ctx, false, errors.New("ip address already exist in "+roomName), nil, http.StatusInternalServerError)
		return
	}

	if err := inputvalid.CheckDeviceInput(&registeredDevice); err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusBadRequest)
		return
	}

	devStatus := device.QueryDevStatus("http://" + registeredDevice.Ipaddr + "/json")

	err = database.Dbman.AddDevice(registeredDevice, devStatus, roomName)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	sendResultJson(ctx, true, nil, nil, http.StatusOK)
}

func showDevices(ctx *gin.Context) {
	roomName := ctx.Param("roomname")

	// check if roomName exist first
	b, err := database.Dbman.CheckRoomExist(roomName)
	if err != nil || !b {
		sendResultJson(ctx, false, errors.New("invalid roomname"), nil, http.StatusInternalServerError)
		return
	}

	devList, devStatus, err := database.Dbman.GetDevices(roomName)

	go device.QueryRoomDevices(devList, devStatus, roomName)

	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	sendResultJson(ctx, true, nil, gin.H{
		"devList":   devList,
		"devStatus": devStatus,
	}, http.StatusOK)
}

func toggleDevice(ctx *gin.Context) {
	roomName := ctx.Param("roomname")
	ipAddr := ctx.Param("ip")
	toggle := ctx.Param("toggle")

	_, devInfo, devStatus, err := database.Dbman.GetDevice(roomName, ipAddr)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	if devInfo.Type == "wled" && devStatus.On_state != (toggle == "on") {
		var wledSwitch wled.WledSwitch
		wledSwitch.On = toggle == "on"
		marshalled, err := json.Marshal(wledSwitch)
		if err != nil {
			sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
			return
		}

		resp, err := http.Post("http://"+ipAddr+"/json", "application/json", bytes.NewBuffer(marshalled))
		if err != nil {
			sendResultJson(ctx, false, err, nil, http.StatusOK)
			return
		}

		fmt.Println(resp)
		// need to check if resp is success also

		defer resp.Body.Close()
	}

	if devStatus.On_state != (toggle == "on") {
		devStatus.On_state = !devStatus.On_state
	}

	err = database.Dbman.UpdateDevStatus(roomName, ipAddr, devStatus)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	sendResultJson(ctx, true, nil, nil, http.StatusOK)
}

func getWledConfigs(ctx *gin.Context) {
	var wledConfig wled.WledConfig

	ip := ctx.Param("ip")

	b, err := database.Dbman.CheckIpExist(ip)
	if err != nil || !b {
		sendResultJson(ctx, false, errors.New("ip address does not exist"), nil, http.StatusInternalServerError)
		return
	}

	client := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get("http://" + ip + "/json")
	if err != nil {
		sendResultJson(ctx, false, errors.New("wled device is offline"), nil, http.StatusOK)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	json.Unmarshal(body, &wledConfig)

	sendResultJson(ctx, true, nil, wledConfig.State, http.StatusOK)
}

func setWled(ctx *gin.Context) {
	// send json to esp device

	ip := ctx.Param("ip")

	var wledState wled.State

	err := ctx.BindJSON(&wledState)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	marshalled, err := json.Marshal(wledState)
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://"+ip+"/json", "application/json", bytes.NewBuffer(marshalled))
	if err != nil {
		sendResultJson(ctx, false, err, nil, http.StatusOK)
		return
	}

	defer resp.Body.Close()

	sendResultJson(ctx, true, nil, nil, http.StatusOK)
}

// func discoverNetworkDevices(ctx *gin.Context) {
// scan for network devices
// using ssdp
// }
