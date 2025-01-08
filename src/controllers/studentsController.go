package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"agenda-backend/src/models"
	"agenda-backend/src/services"
	"agenda-backend/src/utils"
)

func CadastrarAluno(c *gin.Context) {
	var aluno models.Students

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
	var id uint
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	if err := services.SoftDeleteAlunoService(id); err != nil {
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