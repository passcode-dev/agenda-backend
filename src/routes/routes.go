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
			students.POST("", controllers.CadastrarAluno)

			// Deletar estudante
			students.DELETE("", controllers.DeletarAluno)

			// Listar estudantes com paginação
			// GET /students (paginacao, sem id)
			students.GET("", controllers.GetStudents)
		}

		// Rotas relacionadas aos professores
		teachers := api.Group("/teachers")
		{
			// Criar professor
			/*
				{
					"name": "Carlos Oliveira",
					"cpf": "98765412100",
					"birth_date": "1980-01-15"
				}
			*/
			teachers.POST("", controllers.CreateTeacher)

			// Deletar professor
			/*
				{
					"id": 1
				}
			*/
			teachers.DELETE("", controllers.SoftDeleteTeacher)

			// Listar professores com paginação
			///teachers?page=1
			teachers.GET("", controllers.GetAllTeachers)
		}
	}

	return router
}
