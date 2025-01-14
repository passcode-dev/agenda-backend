package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"agenda-backend/src/services"
	"agenda-backend/src/utils"
	"strconv"
	"agenda-backend/src/models"
)

func DeleteTeacherSubjects(c *gin.Context) {
	var payload struct {
		TeacherID  uint   `json:"teacher_id"`
		SubjectIDs []uint `json:"subject_ids"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	if len(payload.SubjectIDs) == 1 {
		// Deleta uma única associação
		if err := services.DeleteTeacherSubjectService(payload.TeacherID, payload.SubjectIDs[0]); err != nil {
			c.JSON(http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to delete association",
				Data:    gin.H{"details": err.Error()},
			})
			return
		}
	} else {
		// Deleta múltiplas associações
		if err := services.DeleteMultipleTeacherSubjectsService(payload.TeacherID, payload.SubjectIDs); err != nil {
			c.JSON(http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to delete associations",
				Data:    gin.H{"details": err.Error()},
			})
			return
		}
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Associations deleted successfully",
	})
}

func DeleteSubject(c *gin.Context) {
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

	if err := services.DeleteSubjectService(payload.ID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to delete subject",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Subject deleted successfully",
	})
}

func UpdateSubject(c *gin.Context) {
	var payload struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	if err := services.UpdateSubjectService(payload.ID, payload.Name); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to update subject",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Subject updated successfully",
	})
}

func GetAllSubjects(c *gin.Context) {
	pageQuery := c.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil || page < 1 {
		page = 1
	}

	subjects, err := services.GetAllSubjectsService(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to retrieve subjects",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Subjects retrieved successfully",
		Data:    subjects,
	})
}

func CreateSubject(c *gin.Context) {
	var payload struct {
		Name      string `json:"name"`
		TeacherID *uint  `json:"teacher_id,omitempty"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	newSubject := models.Subjects{Name: payload.Name}
	if err := services.CreateSubjectService(&newSubject, payload.TeacherID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to create subject",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Status:  "success",
		Message: "Subject created successfully",
		Data:    newSubject,
	})
}

func AssociateTeacherToSubject(c *gin.Context) {
	var payload struct {
		TeacherID uint `json:"teacher_id"`
		SubjectID uint `json:"subject_id"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	if err := services.AssociateTeacherToSubjectService(payload.TeacherID, payload.SubjectID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to associate teacher to subject",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Teacher associated to subject successfully",
	})
}


