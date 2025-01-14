package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"agenda-backend/src/models"
	"agenda-backend/src/services"
	"agenda-backend/src/utils"
)

// CadastrarAluno cria um novo aluno.
// @Summary      Create a new student
// @Description  Cria um novo aluno no sistema.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        student  body   models.Students  true  "Dados do aluno"
// @Success      201   {object} utils.SuccessResponse{status=string,message=string,data=models.Students}
// @Failure      400   {object} utils.ErrorResponse{status=string,message=string}
// @Failure      409   {object} utils.ErrorResponse{status=string,message=string}
// @Failure      500   {object} utils.ErrorResponse{status=string,message=string}
// @Router       /students [post]
func CadastrarAluno(c *gin.Context) {
	var aluno models.Students

	log.Printf("cadastro")

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	if err := services.CreateAlunoService(&aluno); err != nil {
		if err.Error() == "CPF already in use" {
			c.JSON(http.StatusConflict, utils.Response{
				Status:  "error",
				Message: err.Error(),
				Data:    gin.H{"details": err.Error()},
			})
		} else {
			c.JSON(http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to create aluno",
				Data:    gin.H{"details": err.Error()},
			})
		}
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Status:  "success",
		Message: "Aluno cadastrado com sucesso",
		Data:    aluno,
	})
}
// DeletarAluno realiza a exclusão lógica de um aluno.
// @Summary      Soft delete a student
// @Description  Exclui logicamente um aluno no sistema.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        id  body   uint  true  "ID do aluno"
// @Success      200   {object} utils.SuccessResponse{status=string,message=string}
// @Failure      400   {object} utils.ErrorResponse{status=string,message=string}
// @Failure      500   {object} utils.ErrorResponse{status=string,message=string}
// @Router       /students/delete [post]
func DeletarAluno(c *gin.Context) {
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

	if err := services.SoftDeleteAlunoService(payload.ID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to delete aluno",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Aluno deleted successfully",
	})
}
// GetStudents busca todos os alunos com base nos parâmetros de filtro.
// @Summary      Get all students
// @Description  Busca todos os alunos com filtros opcionais (id, nome, rg, cpf, telefone).
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        id    query   string  false  "ID do aluno"
// @Param        name  query   string  false  "Nome do aluno"
// @Param        rg    query   string  false  "RG do aluno"
// @Param        cpf   query   string  false  "CPF do aluno"
// @Param        phone query   string  false  "Telefone do aluno"
// @Param        page  query   int     false  "Página de resultados" default(1)
// @Success      200   {object} utils.SuccessResponse{status=string,message=string,data=[]models.Students}
// @Failure      400   {object} utils.ErrorResponse{status=string,message=string}
// @Failure      500   {object} utils.ErrorResponse{status=string,message=string}
// @Router       /students [get]
func GetStudents(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	name := c.DefaultQuery("name", "")
	rg := c.DefaultQuery("rg", "")
	cpf := c.DefaultQuery("cpf", "")
	phone := c.DefaultQuery("phone", "")

	pageQuery := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageQuery)

	students, err := services.GetAllStudents(id, name, rg, cpf, phone, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to retrieve students",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Students retrieved successfully",
		Data:    students,
	})
}
// UpdateStudent atualiza os dados de um aluno.
// @Summary      Update student data
// @Description  Atualiza os dados de um aluno no sistema.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        student  body   models.Students  true  "Dados do aluno"
// @Success      200   {object} utils.SuccessResponse{status=string,message=string}
// @Failure      400   {object} utils.ErrorResponse{status=string,message=string}
// @Failure      500   {object} utils.ErrorResponse{status=string,message=string}
// @Router       /students/update [put]
func UpdateStudent(c *gin.Context) {
	var payload struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		RG          string `json:"rg"`
		CPF         string `json:"cpf"`
		BirthDate   string `json:"birth_date"`
		PhoneNumber string `json:"phone_number"`
	}

	// Faz o binding do JSON para a estrutura
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	// Monta os dados a serem atualizados
	updatedData := make(map[string]interface{})
	if payload.Name != "" {
		updatedData["name"] = payload.Name
	}
	if payload.RG != "" {
		updatedData["rg"] = payload.RG
	}
	if payload.CPF != "" {
		updatedData["cpf"] = payload.CPF
	}
	if payload.BirthDate != "" {
		updatedData["birth_date"] = payload.BirthDate
	}
	if payload.PhoneNumber != "" {
		updatedData["phone_number"] = payload.PhoneNumber
	}

	// Chama o serviço para atualizar o estudante
	if err := services.UpdateStudentService(payload.ID, updatedData); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to update student",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	// Retorna a resposta de sucesso
	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Student updated successfully",
	})
}
