package socket

import (
	"encoding/json"
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

type onStatusQuery struct {
	RoomName string `json:"roomname"`
	Deviceip string `json:"ipaddr"`
}

func InitSocket() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(c socketio.Conn) error {
		c.SetContext("")
		fmt.Println("connected:", c.ID())
		return nil
	})

	server.OnEvent("/", "testping", func(c socketio.Conn, msg string) {
		fmt.Println("receieved socket msg from client")

		fmt.Println(msg)
	})

	server.OnEvent("/", "devStatus", func(c socketio.Conn, msg string) {
		fmt.Println("receieved socket msg from client")

		var query onStatusQuery

		json.Unmarshal([]byte(msg), &query)

		// devInfo := globalinfo.ServerInfo.Rooms[query.RoomName].DeviceInfo
		// server.BroadcastToNamespace("/", "recvStatus", devInfo)
	})

	return server
}
