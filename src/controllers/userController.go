package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

var validate = validator.New()


func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	created, err := models.CreatedUser(&user)
	if err != nil {
		if err.Error() == "email already in use" {
			c.JSON(http.StatusConflict, Response{
				Status:  "error",
				Message: err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, Response{
				Status:  "error",
				Message: "Failed to create user",
				Data:    gin.H{"details": err.Error()},
			})
		}
		return
	}

	if created {
		c.JSON(http.StatusCreated, Response{
			Status:  "success",
			Message: "User created successfully",
			Data:    user,
		})
	}
}
