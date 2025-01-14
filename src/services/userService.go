package services

import (
	"log"

	"agenda-backend/src/models"
	"agenda-backend/src/repository"
	"agenda-backend/src/utils"
)

func GetUser(email string, id int, username string) ([]models.User, error) {
    return repository.GetUser(email, id, username)
}

func UpdateUser(user *models.UserUpdateRequest, id uint) (error) {
    hashedPassword ,err := utils.HashPassword(user.Password)

    if err != nil {
        return  err
    }

    user.Password = hashedPassword
    log.Print("user.Password ", user.Password)
    return repository.UpdateUser(user, id)
}

func CreatedUser(user *models.UserCreate) (bool, error) {
    hashedPassword,err := utils.HashPassword(user.Password) 

    if err != nil {
        return false, err
    }
    user.Password = hashedPassword

    err = repository.CreateUser(user) 
    if err != nil {
        if err.Error() == "email already in use" {
            return false, err
        }
        return false, err
    }

    return true, nil
}