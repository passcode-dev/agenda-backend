package main

import (
	"mvc/src/routes"
	"mvc/src/database"

)

func main() {
	database.InitDB()
    router := routes.SetupRouter()
	router.Run(":8080")
}
