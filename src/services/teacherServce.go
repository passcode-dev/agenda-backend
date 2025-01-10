package services

import (
	"errors"
	"agenda-backend/src/models"
	"agenda-backend/src/repository"
	"agenda-backend/src/utils"
)

func GetAllTeachersService(page int) ([]models.Teachers, error) {
	return repository.GetAllTeachers(page)
}

func GetTeacherByIDService(id uint) (*models.Teachers, error) {
	return repository.GetTeacherByID(id)
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
	return repository.SoftDeleteTeacher(id)
}