package services

import (
	"errors"
	"agenda-backend/src/models"
	"agenda-backend/src/repository"
	"log" // Adicione esta importação
)

func ValidateAlunoForCreation(aluno *models.Students) error {
	/*if aluno.Name == "" || aluno.PhoneNumber == "" {
		return errors.New("required fields are missing")
	}*/
	log.Printf("JSON recebido:")
	return nil
}

func CreateAlunoService(aluno *models.Students) error {
	log.Printf("entrei")
	if err := ValidateAlunoForCreation(aluno); err != nil {
		return err
	}

	exists, err := repository.VerifyStudentCreated(aluno.CPF)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("CPF already in use")
	}

	return repository.CreateStudent(aluno)
}

func SoftDeleteAlunoService(id uint) error {
	return repository.UpdateDeletedAt(id)
}