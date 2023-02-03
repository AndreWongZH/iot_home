package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status", getServerStatus)
	r.GET("/discover", discoverNetworkDevices)
	r.POST("/add_device", addDevice)
	r.GET("/devices", showDevices)

	return r
}
