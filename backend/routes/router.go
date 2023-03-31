package routes

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/AndreWongZH/iothome/globals"
	"github.com/AndreWongZH/iothome/nmap"
	"github.com/AndreWongZH/iothome/socket"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter(origin string) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{fmt.Sprintf("http://%s", origin)},
		AllowMethods:  []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Set-Cookie"},
		ExposeHeaders: []string{"Content-Length", "Content-Type", "Set-Cookie", "Access-Control-Allow-Credentials", "Access-Control-Expose-Headers", "Access-Control-Allow-Origin", "set-cookie"},

		AllowCredentials: true,
		AllowWebSockets:  true,
		MaxAge:           12 * time.Hour,
	}))

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: false,
		Secure:   false,
	})

	r.Use(sessions.Sessions("mysession", store))

	public := r.Group("/")
	publicRoutes(public)

	private := r.Group("/")
	private.Use(AuthRequired)
	privateRoutes(private)

	private.GET("/ws", socket.WebsocketHandler)

	return r
}

func publicRoutes(g *gin.RouterGroup) {
	g.GET("/", getServerStatus)
	g.POST("/login", loginPost)
	g.POST("/register", registerPost)
}

func privateRoutes(g *gin.RouterGroup) {
	g.GET("/discover", nmap.DiscoverNetworkDevices)

	g.POST("/create-room", createRoom)
	g.GET("/rooms", getRooms)
	g.DELETE("/delete-room/:roomname", deleteRoom)

	g.POST("/:roomname/add-device", addDevice)
	g.GET("/:roomname/devices", showDevices)

	g.GET("/:roomname/:ip", getDeviceInfo)
	g.DELETE("/:roomname/:ip/delete-device", deleteDevice)
	g.POST("/:roomname/:ip/:toggle", toggleDevice)

	g.GET("/:roomname/:ip/wled-config", getWledConfigs)
	g.POST("/:roomname/:ip/wled-config", setWled)

	g.POST("/logout", logoutPost)
	g.GET("/user", getUsername)
}

func AuthRequired(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(globals.UserKey)
	if user == nil {
		sendResultJson(ctx, false, errors.New("user not logged in"), nil, http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
