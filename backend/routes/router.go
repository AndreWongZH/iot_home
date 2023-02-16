package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/status", getServerStatus)

	r.POST("/createrm", createRoom)
	r.GET("/rooms", getRooms)

	r.POST("/:roomname/add_device", addDevice)
	r.GET("/:roomname/devices", showDevices)

	r.GET("/:roomname/wled_config/:ip", getWledConfigs)
	r.POST("/:roomname/wled_config/set/:ip", setWled)

	r.GET("/discover", discoverNetworkDevices)

	return r
}
