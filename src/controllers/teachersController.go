package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"agenda-backend/src/models"
	"agenda-backend/src/services"
	"agenda-backend/src/utils"
)

// GetAllTeachers busca todos os professores com base nos parâmetros de filtro.
// @Summary      Get all teachers
// @Description  Busca todos os professores com filtros opcionais (id, nome, cpf, email, telefone).
// @Tags         Teachers
// @Accept       json
// @Produce      json
// @Param        id    query   int    false  "ID do professor"
// @Param        name  query   string false  "Nome do professor"
// @Param        cpf   query   string false  "CPF do professor"
// @Param        email query   string false  "Email do professor"
// @Param        phone query   string false  "Telefone do professor"
// @Param        page  query   int    false  "Página de resultados" default(1)
// @Success      200   {object} utils.SuccessResponse{status=string,message=string,data=[]models.TeachersResponse}
// @Failure      400   {object} utils.ErrorResponse{status=string,message=string}
// @Failure      500   {object} utils.ErrorResponse{status=string,message=string}
// @Router       /teachers [get]
func GetAllTeachers(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	name := c.DefaultQuery("name", "")
	cpf := c.DefaultQuery("cpf", "")
	email := c.DefaultQuery("email", "")
	phone := c.DefaultQuery("phone", "")

	var teacherID int
	if id != "" {
		var err error
		teacherID, err = strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid ID format",
				Data:    gin.H{"details": err.Error()},
			})
			return
		}
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid page number",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	teachers, err := services.GetAllTeachersService(teacherID, name, cpf, email, phone, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to retrieve teachers",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	// Retorna a lista de professores
	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Teachers retrieved successfully",
		Data:    teachers,
	})
}

// CreateTeacher cria um novo professor.
// @Summary      Create a new teacher
// @Description  Cria um novo professor no sistema.
// @Tags         Teachers
// @Accept       json
// @Produce      json
// @Param        teacher  body   models.Teachers  true  "Dados do professor"
// @Success      201   {object} utils.SuccessResponse{status=string,message=string,data=models.Teachers}
// @Failure      400   {object} utils.ErrorResponse{status=string,message=string}
// @Failure      500   {object} utils.ErrorResponse{status=string,message=string}
// @Router       /teachers [post]
func CreateTeacher(c *gin.Context) {
	var teacher models.Teachers

	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	if err := services.CreateTeacherService(&teacher); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to create teacher",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Status:  "success",
		Message: "Teacher created successfully",
		Data:    teacher,
	})
}

// SoftDeleteTeacher realiza a exclusão lógica de um professor.
// @Summary      Soft delete a teacher
// @Description  Exclui logicamente um professor no sistema.
// @Tags         Teachers
// @Accept       json
// @Produce      json
// @Param        id  body   uint  true  "ID do professor"
// @Success      200   {object} utils.SuccessResponse{status=string,message=string}
// @Failure      400   {object} utils.ErrorResponse{status=string,message=string}
// @Failure      500   {object} utils.ErrorResponse{status=string,message=string}
// @Router       /teachers/delete [post]
func SoftDeleteTeacher(c *gin.Context) {
	var payload struct {
		ID uint `json:"id"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	if err := services.SoftDeleteTeacherService(payload.ID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to delete teacher",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Teacher deleted successfully",
	})
}
