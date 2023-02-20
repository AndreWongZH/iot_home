package main

import (
	"github.com/AndreWongZH/iothome/globalinfo"
	"github.com/AndreWongZH/iothome/routes"
)

func main() {
	globalinfo.InitializeGlobals()

	// socketServer := socket.InitSocket()

	// go socketServer.Serve()
	// defer socketServer.Close()

	r := routes.InitRouter()

	r.Run("localhost:3001")
}
