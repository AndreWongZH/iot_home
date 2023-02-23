package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/globalinfo"
	"github.com/AndreWongZH/iothome/models"
	"github.com/AndreWongZH/iothome/wled"
	"github.com/gin-gonic/gin"
)

func getServerStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func createRoom(ctx *gin.Context) {
	var room models.RoomInfo

	err := ctx.BindJSON(&room)
	if err != nil {
		log.Println("failed to bind to json")
		log.Println(err)
		return
	}

	err = database.Dbman.AddRoom(room)
	if err != nil {
		log.Println(err)
	}
}

func getRooms(ctx *gin.Context) {
	rooms, err := database.Dbman.GetRooms()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rooms,
	})
}

func addDevice(ctx *gin.Context) {
	var registeredDevice models.RegisteredDevice

	roomName := ctx.Param("roomname")

	err := ctx.BindJSON(&registeredDevice)
	if err != nil {
		log.Println("error binding json data to variable")
	}

	err = database.Dbman.AddDevice(registeredDevice, roomName)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, gin.H{"success": false, "error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func showDevices(ctx *gin.Context) {
	roomName := ctx.Param("roomname")

	devList, devStatus, err := database.Dbman.GetDevices(roomName)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":   true,
		"devList":   devList,
		"devStatus": devStatus,
	})
}

func offDevice(ctx *gin.Context) {
	roomName := ctx.Param("roomname")
	ip := ctx.Param("ip")

	var wledSwitch wled.WledSwitch
	wledSwitch.On = false

	marshalled, err := json.Marshal(wledSwitch)
	if err != nil {
		log.Println("error marshalling data")
	}

	if devStatus, ok := globalinfo.ServerInfo.Rooms[roomName].DeviceInfo[ip]; ok {
		if deviceType, ok := globalinfo.ServerInfo.Rooms[roomName].Devices[ip]; ok {
			if deviceType.Type == "wled" {
				resp, err := http.Post("http://"+ip+"/json", "application/json", bytes.NewBuffer(marshalled))
				if err != nil {
					log.Println("error")
					log.Println(err)

					ctx.JSON(http.StatusOK, gin.H{
						"success": false,
						"error":   "failed to post to device",
					})
				}

				fmt.Println(resp)
				// need to check if resp is success also

				defer resp.Body.Close()
			}

			devStatus.On = false
			globalinfo.ServerInfo.Rooms[roomName].DeviceInfo[ip] = devStatus

			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"error":   "roomname or ip address not found",
	})
}

func onDevice(ctx *gin.Context) {
	roomName := ctx.Param("roomname")
	ip := ctx.Param("ip")

	var wledSwitch wled.WledSwitch
	wledSwitch.On = true

	marshalled, err := json.Marshal(wledSwitch)
	if err != nil {
		log.Println("error marshalling data")
	}

	devStatus := globalinfo.ServerInfo.Rooms[roomName].DeviceInfo[ip]

	if globalinfo.ServerInfo.Rooms[roomName].Devices[ip].Type == "wled" {
		resp, err := http.Post("http://"+ip+"/json", "application/json", bytes.NewBuffer(marshalled))
		if err != nil {
			log.Println("error")
			log.Println(err)
		}

		fmt.Println(resp)
		defer resp.Body.Close()
	}

	devStatus.On = true
	globalinfo.ServerInfo.Rooms[roomName].DeviceInfo[ip] = devStatus

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func getWledConfigs(ctx *gin.Context) {
	var wledConfig wled.WledConfig

	ip := ctx.Param("ip")

	fmt.Println("ip addr to find config is:", ip)

	client := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get("http://" + ip + "/json")
	if err != nil {
		log.Println("error retrieving wled configs")
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "error retrieving wled configs",
		})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading wled configs")
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "error reading wled configs",
		})
		return
	}

	json.Unmarshal(body, &wledConfig)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    wledConfig.State,
		"success": true,
	})
}

func setWled(ctx *gin.Context) {
	// send json to esp device

	ip := ctx.Param("ip")

	var wledState wled.State

	err := ctx.BindJSON(&wledState)
	if err != nil {
		log.Println("error binding to wled state")
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "error binding wled state",
		})
	}

	marshalled, err := json.Marshal(wledState)
	if err != nil {
		log.Println("error marshalling data")
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "error marshalling data",
		})
	}
	fmt.Println(wledState)

	resp, err := http.Post("http://"+ip+"/json", "application/json", bytes.NewBuffer(marshalled))
	if err != nil {
		log.Println("error posting to device ip")
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "error posting to device ip",
		})
	}

	defer resp.Body.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func discoverNetworkDevices(ctx *gin.Context) {
	// scan for network devices
	// using ssdp
}
