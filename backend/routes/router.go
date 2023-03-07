package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AndreWongZH/iothome/globals"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// func GinMiddleware(allowOrigin string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 	}
// }

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// func InitRouter(socketServer *socketio.Server) *gin.Engine {
func InitRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(GinMiddleware("http://localhost:3000"))

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"POST", "GET", "PUT", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Set-Cookie"},
		ExposeHeaders: []string{"Content-Length", "Content-Type", "Set-Cookie", "Access-Control-Allow-Credentials", "Access-Control-Expose-Headers", "Access-Control-Allow-Origin", "set-cookie"},

		AllowCredentials: true,
		AllowWebSockets:  true,
		MaxAge:           12 * time.Hour,
	}))

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * 24 * 7,
		// SameSite: http.SameSiteNoneMode,
		HttpOnly: false,
		Secure:   false,
	})

	r.Use(sessions.Sessions("mysession", store))

	public := r.Group("/")
	publicRoutes(public)

	private := r.Group("/")
	private.Use(AuthRequired)
	privateRoutes(private)

	private.GET("/ws", func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println("failed to upgrade to websockets")
		}
		fmt.Println("client is connected")
		defer conn.Close()

		type Data struct {
			Success string `json:"success"`
		}

		var datapacket Data

		for {
			// messageType, p, err := conn.ReadJSON(&datapacket)
			err := conn.ReadJSON(&datapacket)
			if err != nil {
				fmt.Println(err)
				return
			}
			// fmt.Println("message type: ", messageType)
			// fmt.Println("data: ", p)
			fmt.Println("data: ", datapacket.Success)

			// err = conn.WriteMessage(messageType, p)
			// if err != nil {
			// 	return
			// }
		}
	})

	// r.GET("/discover", discoverNetworkDevices)

	// r.GET("/socket.io/*any", gin.WrapH(socketServer))
	// r.POST("/socket.io/*any", gin.WrapH(socketServer))

	return r
}

func publicRoutes(g *gin.RouterGroup) {
	g.GET("/", getServerStatus)
	g.POST("/login", loginPost)
}

func privateRoutes(g *gin.RouterGroup) {
	g.POST("/createrm", createRoom)
	g.GET("/rooms", getRooms)

	g.POST("/:roomname/add_device", addDevice)
	g.GET("/:roomname/devices", showDevices)

	g.POST("/:roomname/:ip/:toggle", toggleDevice)

	g.GET("/:roomname/wled_config/:ip", getWledConfigs)
	g.POST("/:roomname/wled_config/set/:ip", setWled)
}

func AuthRequired(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(globals.UserKey)
	// fmt.Println(ctx.Request.Header)
	// fmt.Println("user sessions:", user)
	if user == nil {
		sendResultJson(ctx, false, errors.New("user not logged in"), nil)
		return
	}

	ctx.Next()
}
