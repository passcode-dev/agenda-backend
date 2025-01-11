package services

import (
	"agenda-backend/src/models"
	"agenda-backend/src/repository"
	"agenda-backend/src/utils"
	"errors"
)

func GetAllTeachersService(id int, name, cpf, email, phone string, page int) ([]models.Teachers, error) {
	return repository.GetAllTeachers(id, name, cpf, email, phone, page)
}

func CreateTeacherService(teacher *models.Teachers) error {
	if teacher.Name == "" || teacher.CPF == "" || teacher.BirthDate == "" {
		return errors.New("Todos os campos obrigatórios devem ser preenchidos")
	}

	if err := utils.ValidateCPF(teacher.CPF); err != nil {
		return err
	}

	return repository.CreateTeacher(teacher)
}

func SoftDeleteTeacherService(id uint) error {
	// Chama o repositório para realizar o soft delete
	return repository.SoftDeleteTeacher(id)
}
