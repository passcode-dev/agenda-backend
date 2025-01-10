package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"agenda-backend/src/services"
	"agenda-backend/src/utils"
	"agenda-backend/src/models"
)

func GetAllTeachers(c *gin.Context) {
	pageQuery := c.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil || page < 1 {
		page = 1
	}

	teachers, err := services.GetAllTeachersService(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to retrieve teachers",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Teachers retrieved successfully",
		Data:    teachers,
	})
}

func GetTeacherByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid ID",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	teacher, err := services.GetTeacherByIDService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to retrieve teacher",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Teacher retrieved successfully",
		Data:    teacher,
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