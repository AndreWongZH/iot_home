package main

import (
	"fmt"
	"os"

	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/device"
	"github.com/AndreWongZH/iothome/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	db := database.InitDatabase()
	database.InitializeGlobals(db)

	r := routes.InitRouter()

	exit := make(chan bool)
	go device.QueryAllDevices(exit)

	fmt.Println("Starting server on port", port)
	r.Run(":" + port)

	defer cleanup(exit)
}

func cleanup(exit chan bool) {
	exit <- true
}
