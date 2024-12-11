package main

import (

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	/*client, _ := mongo.NewClient()
	repo := repository.NewUserRepository(client)
	control := controller.NewUserController(repo)*/
	config.InitDB()
    router := routes.SetupRouter()
	router.Run(":8080")
}
