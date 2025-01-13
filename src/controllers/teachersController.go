package controllers

import (
	"net/http"
	"strconv"

	"agenda-backend/src/models"
	"agenda-backend/src/services"
	"agenda-backend/src/utils"

	"github.com/gin-gonic/gin"
)

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
