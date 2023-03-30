package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndreWongZH/iothome/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type DevStatusesMsg struct {
	DevStatuses map[string]models.DeviceStatus `json:"devstatuses"`
	RoomName    string                         `json:"roomname"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connectedSockets []*websocket.Conn

func WebsocketHandler(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("failed to upgrade to websockets")
	}

	// if ws conn is closed, remove from our array of connectedSockets
	conn.SetCloseHandler(func(code int, text string) error {
		for i, c := range connectedSockets {
			if c == conn {
				connectedSockets = append(connectedSockets[:i], connectedSockets[i+1:]...)
				break
			}
		}
		fmt.Println("websocket is closed")
		return nil
	})

	fmt.Println("client is connected")
	defer conn.Close()
	connectedSockets = append(connectedSockets, conn)

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
	}
}

func BroadcastMsg(data interface{}) {
	for _, conn := range connectedSockets {
		conn.WriteJSON(data)
	}
}
