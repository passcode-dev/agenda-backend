// src/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "agenda-backend/docs" // Importa os documentos gerados
	"agenda-backend/src/controllers"
	"agenda-backend/src/utils"
)

// SetupRouter configura todas as rotas da API
// @title           Agenda Backend API
// @version         1.0
// @description     API para gerenciamento de usuários, estudantes e professores.
// @termsOfService  http://example.com/terms/

// @contact.name   Suporte
// @contact.url    http://example.com/support
// @contact.email  suporte@example.com

// @host      localhost:8080
// @BasePath  /api
func SetupRouter() *gin.Engine {
	// Inicializa o router
	router := gin.Default()

	// Middleware para CORS
	router.Use(utils.CORSMiddleware())

	// Rota de teste
	// @Summary      Testar conexão
	// @Description  Testa a conexão com o servidor
	// @Tags         Test
	// @Produce      json
	// @Success      200  {object}  map[string]string
	// @Router       / [get]
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testando",
		})
	})

	// Rotas da API principal
	api := router.Group("/api")
	{
		// Rotas relacionadas à autenticação e usuários
		// @Summary      Login de usuário
		// @Description  Realiza login do usuário
		// @Tags         Auth
		// @Accept       json
		// @Produce      json
		// @Param        login  body      controllers.LoginRequest  true  "Dados de login"
		// @Success      200    {object} controllers.LoginResponse
		// @Failure      401    {object} utils.ErrorResponse
		// @Router       /login [post]
		api.POST("/login", controllers.AuthLogin)

		user := api.Group("/user")
		{
			// @Summary      Criar usuário
			// @Description  Cadastra um novo usuário
			// @Tags         User
			// @Accept       json
			// @Produce      json
			// @Param        user  body      controllers.CreateUserRequest  true  "Dados do usuário"
			// @Success      201   {object} controllers.UserResponse
			// @Router       /user [post]
			user.POST("", controllers.CreateUser)

			// @Summary      Obter usuário
			// @Description  Retorna os dados de um usuário
			// @Tags         User
			// @Produce      json
			// @Param        email     query     string  false  "Email do usuário"
			// @Param        id        query     int     false  "ID do usuário"
			// @Param        username  query     string  false  "Nome de usuário"
			// @Success      200       {array}   controllers.UserResponse
			// @Router       /user [get]
			user.GET("", controllers.GetUser)
			
			// @Summary      Alterar usuário
			
			user.PUT(":id", controllers.UpdateUser)

		}

		students := api.Group("/students")
		{
			// @Summary      Criar estudante
			// @Description  Cadastra um novo estudante
			// @Tags         Students
			// @Accept       json
			// @Produce      json
			// @Param        student  body      controllers.StudentRequest  true  "Dados do estudante"
			// @Success      201      {object} controllers.StudentResponse
			// @Router       /students [post]
			students.POST("", controllers.CadastrarAluno)

			// @Summary      Deletar estudante
			// @Description  Remove um estudante pelo ID
			// @Tags         Students
			// @Accept       json
			// @Produce      json
			// @Param        id  body      int  true  "ID do estudante"
			// @Success      204  {string} string  "No Content"
			// @Router       /students [delete]
			students.DELETE("", controllers.DeletarAluno)

			// @Summary      Listar estudantes
			// @Description  Retorna uma lista de estudantes com paginação
			// @Tags         Students
			// @Produce      json
			// @Param        page  query     int  false  "Número da página"
			// @Success      200   {array}   controllers.StudentResponse
			// @Router       /students [get]
			students.GET("", controllers.GetStudents)

			students.GET("/:id", controllers.GetStudentID)

		}

		teachers := api.Group("/teachers")
		{
			// @Summary      Criar professor
			// @Description  Cadastra um novo professor
			// @Tags         Teachers
			// @Accept       json
			// @Produce      json
			// @Param        teacher  body      controllers.TeacherRequest  true  "Dados do professor"
			// @Success      201      {object} controllers.TeacherResponse
			// @Router       /teachers [post]
			teachers.POST("", controllers.CreateTeacher)

			// @Summary      Deletar professor
			// @Description  Remove um professor pelo ID
			// @Tags         Teachers
			// @Accept       json
			// @Produce      json
			// @Param        id  body      int  true  "ID do professor"
			// @Success      204  {string} string  "No Content"
			// @Router       /teachers [delete]
			teachers.DELETE("", controllers.SoftDeleteTeacher)

			// @Summary      Listar professores
			// @Description  Retorna uma lista de professores com paginação
			// @Tags         Teachers
			// @Produce      json
			// @Param        page  query     int  false  "Número da página"
			// @Success      200   {array}   controllers.TeacherResponse
			// @Router       /teachers [get]
			teachers.GET("", controllers.GetAllTeachers)
		}

		// Rotas relacionadas às matérias
		subjects := api.Group("/subjects")
		{
			subjects.GET("", controllers.GetAllSubjects)                   // Listar matérias
			subjects.POST("", controllers.CreateSubject)                   // Criar matéria
			subjects.POST("/associate", controllers.AssociateTeacherToSubject) // Associar professor a matéria
			subjects.DELETE("/teacher-subjects", controllers.DeleteTeacherSubjects) // Deletar associações de professores
			subjects.DELETE("", controllers.DeleteSubject)                 // Deletar matéria
			subjects.PUT("", controllers.UpdateSubject)                    // Atualizar matéria
		}
	}
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
