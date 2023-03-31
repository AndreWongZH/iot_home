package main

import (
	"os"

	"github.com/AndreWongZH/iothome/database"
	"github.com/AndreWongZH/iothome/device"
	"github.com/AndreWongZH/iothome/logger"
	"github.com/AndreWongZH/iothome/routes"
)

func getEnvVar() (string, string) {
	port := os.Getenv("PORT")
	origin := os.Getenv("ORIGIN")
	if port == "" {
		port = "3001"
	}
	if origin == "" {
		origin = "localhost:3000"
	}

	logger.SugarLog.Info("origin used: ", origin)
	return origin, port
}

func main() {
	logger.InitLogger()

	origin, port := getEnvVar()

	db := database.InitDatabase()
	database.InitializeGlobals(db)

	r := routes.InitRouter(origin)

	exit := make(chan bool)
	go device.QueryAllDevices(exit)

	logger.SugarLog.Info("Starting server on port: ", port)
	r.Run(":" + port)

	defer cleanup(exit)
}

func cleanup(exit chan bool) {
	exit <- true
}
