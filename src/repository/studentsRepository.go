package repository

import (
	"time"
	"agenda-backend/src/database"
	"agenda-backend/src/models"
	"log"
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



func CreateStudent(student *models.Students) error {
	return database.DB.Create(student).Error
}

func UpdateDeletedAt(id uint) error {
	currentTime := time.Now()
	return database.DB.Model(&models.Students{}).Where("id = ?", id).Update("deleted_at", &currentTime).Error
}

func GetStudents(page int) ([]models.Students, error) {
	var students []models.Students
	offset := (page - 1) * 10 // Calcula o offset com base na página atual

	// Consulta com WHERE deleted_at IS NULL e LIMIT para paginação
	err := database.DB.Where("deleted_at IS NULL").Limit(10).Offset(offset).Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}

func UpdateStudent(id uint, updatedData map[string]interface{}) error {
	// Atualiza os campos especificados com base no ID
	return database.DB.Model(&models.Students{}).Where("id = ?", id).Updates(updatedData).Error
}
