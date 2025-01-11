package repository

import (
	"agenda-backend/src/database"
	"agenda-backend/src/models"
	"log"
	"time"
)

func VerifyStudentCreated(cpf string) (bool, error) {
	var count int64

	// Conta o número de registros com o CPF informado
	err := database.DB.Model(&models.Students{}).
		Where("cpf = ? AND deleted_at IS NULL", cpf).
		Count(&count).Error

	if err != nil {
		log.Printf("Erro ao verificar CPF %s no banco: %s", cpf, err)
		return false, err // Retorna erro caso haja problemas na consulta
	}

	// Verifica se o CPF existe com base na contagem
	if count > 0 {
		log.Printf("CPF %s encontrado no banco", cpf)
		return true, nil // CPF já existe
	}

	log.Printf("CPF %s não encontrado no banco", cpf)
	return false, nil // CPF não existe
}

func GetStudentById(id uint) (*models.Students, error) {
	log.Printf("ID recebido: %d", id)
	var student models.Students
	err := database.DB.Model(&models.Students{}).
		Where("deleted_at IS NULL AND id = ?", id).
		First(&student).Error

	return &student, err
}

func CreateStudent(student *models.Students) error {
	return database.DB.Create(student).Error
}

func UpdateDeletedAt(id uint) error {
	currentTime := time.Now()
	return database.DB.Model(&models.Students{}).Where("id = ?", id).Update("deleted_at", &currentTime).Error
}

func GetStudents(id, name, rg, cpf, phone string, page int) ([]models.Students, error) {
	var students []models.Students

	query := database.DB.Model(&models.Students{})

	if id != "" {
		query = query.Where("id = ?", id)
	}
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if rg != "" {
		query = query.Where("rg = ?", rg)
	}
	if cpf != "" {
		query = query.Where("cpf = ?", cpf)
	}
	if phone != "" {
		query = query.Where("phone = ?", phone)
	}

	limit := 10
	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	err := query.Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}

func UpdateStudent(id uint, updatedData map[string]interface{}) error {
	// Atualiza os campos especificados com base no ID
	return database.DB.Model(&models.Students{}).Where("id = ?", id).Updates(updatedData).Error
}
