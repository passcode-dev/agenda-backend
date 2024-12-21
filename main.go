package main

import (
	"mvc/src/routes"
	"mvc/src/config"

)

func main() {
	config.InitDB()
    router := routes.SetupRouter()
	router.Run(":8080")
}
