package main

import (
	"github.com/AndreWongZH/iothome/globalinfo"
	"github.com/AndreWongZH/iothome/routes"
)

func main() {
	globalinfo.InitializeGlobals()

	r := routes.InitRouter()
	r.Run("localhost:3000")
}
