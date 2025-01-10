package repository

import (
	"agenda-backend/src/database"
	"agenda-backend/src/models"
	"time"
)

func GetAllTeachers(page int) ([]models.Teachers, error) {
	var teachers []models.Teachers
	offset := (page - 1) * 10
	if err := database.DB.Where("deleted_at IS NULL").Limit(10).Offset(offset).Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func GetTeacherByID(id uint) (*models.Teachers, error) {
	var teacher models.Teachers
	if err := database.DB.Where("id = ? AND deleted_at IS NULL", id).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func CreateTeacher(teacher *models.Teachers) error {
	return database.DB.Create(teacher).Error
}

func SoftDeleteTeacher(id uint) error {
	currentTime := time.Now()
	return database.DB.Model(&models.Teachers{}).Where("id = ?", id).Update("deleted_at", &currentTime).Error
}