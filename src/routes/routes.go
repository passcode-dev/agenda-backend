// src/routes/routes.go
package routes

import (
	"agenda-backend/src/controllers"
	"agenda-backend/src/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Inicializa o router
	router := gin.Default()
	
	// Middleware para CORS
	router.Use(utils.CORSMiddleware())

	// Rota de teste
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testando",
		})
	})

	// Rotas da API principal
	api := router.Group("/api")
	{
		// Rotas relacionadas à autenticação e usuários
		api.POST("/login", controllers.AuthLogin)
		api.POST("/users", controllers.CreateUser)

		// Rotas relacionadas aos estudantes
		students := api.Group("/students")
		{
			// Criar estudante
			/*
			{
				"name": "Leonardo Silva",
				"rg": "123456789",
				"cpf": "98765412100",
				"birth_date": "2006-01-02",
				"phone_number": "11987654321"
			}
			*/
			students.POST("", controllers.CadastrarAluno)

			// Deletar estudante
			/*
			{
				"id": 1
			}
			*/
			students.DELETE("", controllers.DeletarAluno)

			// Listar estudantes com paginação
			///students?page=1
			students.GET("", controllers.GetStudents)

			// Atualizar estudante
			/*
			{
				"id": 1,
				"name": "Leonardo Silva",
				"rg": "123456789",
				"cpf": "23541241141144411124100",
				"birth_date": "2006-01-02",
				"phone_number": "11987654321"
			}
			*/
			students.PUT("", controllers.UpdateStudent)
		}
	}

	return router
}
