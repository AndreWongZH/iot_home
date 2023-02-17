package main

import (
	"github.com/AndreWongZH/iothome/globalinfo"
	"github.com/AndreWongZH/iothome/routes"
	"github.com/AndreWongZH/iothome/socket"
)

func main() {
	globalinfo.InitializeGlobals()

	socketServer := socket.InitSocket()

	go socketServer.Serve()
	defer socketServer.Close()

	r := routes.InitRouter(socketServer)

	r.Run("localhost:3001")
}
