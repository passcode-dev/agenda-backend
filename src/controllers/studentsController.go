package controllers

import (
	"log"
	"net/http"

	"agenda-backend/src/models"
	"agenda-backend/src/services"
	"agenda-backend/src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudentsByID(c *gin.Context) {
	idParam := c.Query("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to retrieve students",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	students, err := services.GetStudentsByIDService(uint(id))

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "students retrieved successfully",
		Data:    students,
	})
}

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
