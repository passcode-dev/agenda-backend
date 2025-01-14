package main

import (
	"agenda-backend/src/database"
	"agenda-backend/src/routes"
	"log"
)

func main() {
	log.Println("Inicializando o servidor...")
	database.InitDB()
    router := routes.SetupRouter()
	router.Run(":8080")
}

/*
AJUSTAR UPDATE PARA VALIDAÇÃO DO CPF E MODEL EM VEZ DE MAP STRING
rota update teachers
não dexar cpf dentcos cadastros

colocar data de entrada de alunos, e saida de alunos
*/

