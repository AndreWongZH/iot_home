package main

import (
	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/device"
	"github.com/AndreWongZH/iothome/routes"
)

func main() {
	db := database.InitDatabase()
	database.InitializeGlobals(db)

	// socketServer := socket.InitSocket()

	// go socketServer.Serve()
	// defer socketServer.Close()

	r := routes.InitRouter()

	exit := make(chan bool)
	go device.QueryAllDevices(exit)

	r.Run("localhost:3001")

	defer cleanup(exit)
}

func cleanup(exit chan bool) {
	exit <- true
}
