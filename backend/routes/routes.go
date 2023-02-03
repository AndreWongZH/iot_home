package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AndreWongZH/iothome/globalinfo"
	"github.com/gin-gonic/gin"
)

func getServerStatus(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "ok"})
}

func discoverNetworkDevices(ctx *gin.Context) {
	// scan for network devices
}

func addDevice(ctx *gin.Context) {
	var registeredDevice globalinfo.RegisteredDevice

	err := ctx.BindJSON(&registeredDevice)
	if err != nil {
		log.Println("error binding json data to variable")
	}

	globalinfo.ServerInfo.Devices = append(globalinfo.ServerInfo.Devices, registeredDevice)

	fmt.Println("device :", registeredDevice, "is added")
}

func showDevices(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, globalinfo.ServerInfo.Devices)
}

func getWledConfigs(ctx *gin.Context) {
	// var wledConfig interface{}

	resp, err := http.Get("192.129.23.1/json")
	if err != nil {
		log.Println("error retrieving wled configs")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading wlead configs")
	}

	fmt.Println(string(body))

}

func setWled(ctx *gin.Context) {
	// send json to esp device

	// var wledConfig interface{}

	// err := ctx.BindJSON(&wledConfig)
	// http.Post("192.129.1.1/json", "application/json", wledConfig)

}
