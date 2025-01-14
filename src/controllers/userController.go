package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"agenda-backend/src/models"
	"agenda-backend/src/services"
	"agenda-backend/src/utils"
)

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

var validate = validator.New()

// @Summary      Obter usuário
// @Description  Retorna um usuário com base no email, id ou nome de usuário
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        email    query   string  false  "Email do usuário"
// @Param        username query   string  false  "Nome de usuário"
// @Param        id       query   int     false  "ID do usuário"
// @Success      200      {object} models.User
// @Failure      404      {object} utils.Response  "Usuário não encontrado"
// @Router       /user [get]
func GetUser(c *gin.Context) {
	email := c.DefaultQuery("email","")
	username := c.DefaultQuery("username","")
	id, err := strconv.Atoi(c.DefaultQuery("id",""))
	
	user, err := services.GetUser(email, id, username)
	
	if err != nil {
		c.JSON(http.StatusNotFound, utils.Response{
			Status:  "error",
			Message: "User not found",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "User found",
		Data:    user,
	})
}

func UpdateUser(c *gin.Context) {
	var userUpdateRequest models.UserUpdateRequest

	// Pega o id da URL e converte para uint
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid request data",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	log.Print(userUpdateRequest, id)

	// Passa o ponteiro para a função UpdateUser
	err = services.UpdateUser(&userUpdateRequest, uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to update user",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  "success",
		Message: "User updated successfully",
		Data:    userUpdateRequest, // Retorna os dados atualizados ou qualquer outro objeto que você queira
	})
}


// Criar usuário
// @Summary      Criar usuário
// @Description  Cria um novo usuáio no sistema
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user    body    models.UserCreate  true  "Dados do usuário" 
// @Success      201     {object}  models.UserResponse
// @Router       /user [post]
func CreateUser(c *gin.Context) {
	

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Data:    gin.H{"details": err.Error()},
		})
		return
	}

	created, err := services.CreatedUser(&user)
	if err != nil {
		if err.Error() == "email already in use" {
			c.JSON(http.StatusConflict, utils.Response{
				Status:  "error",
				Message: err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to create user",
				Data:    gin.H{"details": err.Error()},
			})
		}
		return
	}

	if created {
		c.JSON(http.StatusCreated, utils.Response{
			Status:  "success",
			Message: "User created successfully",
			Data:    user,
		})
	}
}
