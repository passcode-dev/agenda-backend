package main

import (
	"agenda-backend/src/routes"
	"agenda-backend/src/database"
)

func main() {
	database.InitDB()
    router := routes.SetupRouter()
	router.Run(":8080")
}

/*
AJUSTAR UPDATE PARA VALIDAÇÃO DO CPF E MODEL EM VEZ DE MAP STRING

