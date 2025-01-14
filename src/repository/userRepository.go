package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"agenda-backend/src/database"
	"agenda-backend/src/models"
)


func VerifyUserCreated(email string) (bool, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func GetUser(email string, id int, username string) ([]models.User, error) {
    var users []models.User

    query := database.DB.Model(&models.User{})
    if email != "" {
        query = query.Where("email = ?", email)
    }
    if id != 0 {
        query = query.Where("id = ?", id)
    }
    if username != "" {
        query = query.Where("username = ?", username)
    }

    err := query.Find(&users).Error
    if err != nil {
        return nil, err
    }

	return users, nil
}

func UpdateUser(user *models.UserUpdateRequest, id uint) error {
	log.Print(user, id)

	return database.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func CreateUser(user *models.User) error {
	
	exists, err := VerifyUserCreated(user.Email)
	if err != nil {
		return err 
	}

	if exists {
		return errors.New("email already in use") 
	}

	if err := database.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}