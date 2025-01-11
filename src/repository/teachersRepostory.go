package repository

import (
	"agenda-backend/src/database"
	"agenda-backend/src/models"
	"time"
)

func GetAllTeachers(id int, name, cpf, email, phone string, page int) ([]models.Teachers, error) {
	var teachers []models.Teachers

	// Calcular o offset para a paginação
	const limit = 10
	offset := (page - 1) * limit

	// Criar uma query base
	query := database.DB.Model(&models.Teachers{})

	// Adicionar filtros à query, se fornecidos
	if id != 0 {
		query = query.Where("id = ?", id)
	}
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if cpf != "" {
		query = query.Where("cpf LIKE ?", "%"+cpf+"%")
	}
	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}
	if phone != "" {
		query = query.Where("phone LIKE ?", "%"+phone+"%")
	}

	// Executa a query com paginação
	if err := query.Offset(offset).Limit(limit).Find(&teachers).Error; err != nil {
		return nil, err
	}

	return teachers, nil
}
func CreateTeacher(teacher *models.Teachers) error {
	return database.DB.Create(teacher).Error
}

func SoftDeleteTeacher(id uint) error {
	currentTime := time.Now()
	return database.DB.Model(&models.Teachers{}).Where("id = ?", id).Update("deleted_at", &currentTime).Error
}
