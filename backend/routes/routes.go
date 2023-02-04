package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AndreWongZH/iothome/globalinfo"
	"github.com/AndreWongZH/iothome/wled"
	"github.com/gin-gonic/gin"
)

func getServerStatus(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "ok"})
}

func discoverNetworkDevices(ctx *gin.Context) {
	// scan for network devices
	// using ssdp
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
		log.Println("error reading wlead configs")
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
