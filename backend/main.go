package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

type Network struct {
	Hostname string `json:"Hostname"`
	Ipaddr   string `json:"Ipaddr"`
}

type RegisteredDevice struct {
	Hostname string `json:"hostname"`
	Ipaddr   string `json:"ipaddr"`
	Nickname string `json:"nickname"`
	Type     string `json:"type"`
}

func initRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/discover", func(ctx *gin.Context) {

		// scan for network devices
		test := []Network{
			{Hostname: "tutu", Ipaddr: "213123"},
		}
		fmt.Println(test)

		ctx.JSON(200, gin.H{
			"andre": "192.168.5.1",
			"magic": []Network{{Hostname: "tutu", Ipaddr: "213123"}, {Hostname: "tutu", Ipaddr: "dud"}},
		})
	})

	r.POST("/add_device", func(ctx *gin.Context) {
		info, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			// handle error here
			log.Println("error")
		}

		fmt.Println(string(info))
		var result RegisteredDevice
		err = json.Unmarshal(info, &result)
		if err != nil {
			log.Println("error parsing the json data")
			log.Println(err)
		}

		fmt.Println(result)

		// ipaddr := ctx.PostForm("ipaddr")
		// hostname := ctx.PostForm("hostname")
		// nickname := ctx.PostForm("nickname")
		// devtype := ctx.PostForm("devtype")
	})

	return r
}

func main() {
	r := initRouter()
	// fmt.Println(r)
	r.Run()
}
