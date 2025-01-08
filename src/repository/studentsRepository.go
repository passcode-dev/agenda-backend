package repository

import (
	"errors"
	"time"
	"agenda-backend/src/database"
	"agenda-backend/src/models"
	"gorm.io/gorm"
)

func VerifyStudentCreated(cpf string) (bool, error) {
	var student models.Students
	err := database.DB.Where("cpf = ?", cpf).First(&student).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return err == nil, err
}

func CreateStudent(student *models.Students) error {
	return database.DB.Create(student).Error
}

func UpdateDeletedAt(id uint) error {
	currentTime := time.Now()
	return database.DB.Model(&models.Students{}).Where("id = ?", id).Update("deleted_at", &currentTime).Error
}
