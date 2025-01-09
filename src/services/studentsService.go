package services

import (
	"agenda-backend/src/models"
	"agenda-backend/src/repository"
	"errors"
	"log"
	"agenda-backend/src/utils"
)

func ValidateAlunoForCreation(aluno *models.Students) error {
	if aluno.Name == "" || aluno.PhoneNumber == "" {
		log.Printf("1")

		return errors.New("required fields are missing")
	}

	if aluno.CPF != ""{

		if err := utils.ValidateCPF(aluno.CPF); err != nil {
			return err
		}

	}
	return nil
}

func CreateAlunoService(aluno *models.Students) error {
	log.Printf("Iniciando criação do aluno com CPF: %s", aluno.CPF)

	// Valida os dados do aluno
	if err := ValidateAlunoForCreation(aluno); err != nil {
		log.Printf("Erro na validação: %s", err)
		return err
	}

	// Verifica se o CPF já existe no banco
	exists, err := repository.VerifyStudentCreated(aluno.CPF)
	if err != nil {
		log.Printf("Erro ao verificar CPF no banco: %s", err)
		return err // Propaga o erro real
	}

	// Se o CPF já existe, retorna erro
	if exists {
		log.Printf("CPF %s já está em uso", aluno.CPF)
		return errors.New("CPF already in use")
	}

	// CPF não existe, cria o aluno
	log.Printf("Criando novo registro para o CPF: %s", aluno.CPF)
	if err := repository.CreateStudent(aluno); err != nil {
		log.Printf("Erro ao criar aluno no banco: %s", err)
		return err
	}

	log.Printf("Aluno criado com sucesso para o CPF: %s", aluno.CPF)
	return nil
}

func SoftDeleteAlunoService(id uint) error {
	return repository.UpdateDeletedAt(id)
}
func GetAllStudents(page int) ([]models.Students, error) {
	// Chama o repositório para obter os estudantes
	return repository.GetStudents(page)
}

func UpdateStudentService(id uint, updatedData map[string]interface{}) error {

	// Verifica se há dados para atualizar
	if len(updatedData) == 0 {
		return errors.New("no fields to update")
	}

	// Atualiza os dados no repositório
	return repository.UpdateStudent(id, updatedData)
}