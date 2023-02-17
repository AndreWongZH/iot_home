package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

func InitRouter(socketServer *socketio.Server) *gin.Engine {
	r := gin.Default()

	// r.Use(GinMiddleware("http://localhost:3000"))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
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

	r.GET("/socket.io/*any", gin.WrapH(socketServer))
	r.POST("/socket.io/*any", gin.WrapH(socketServer))

	return r
}
