package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status", getServerStatus)
	r.GET("/discover", discoverNetworkDevices)
	r.POST("/add_device", addDevice)
	r.GET("/devices", showDevices)

	r.GET("/wled_config/:ip", getWledConfigs)
	r.POST("/wled_config/set/:ip", setWled)

	return r
}
