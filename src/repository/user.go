package repositories

import (
	"errors"
	"mvc/src/database"
	"mvc/src/models"
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


// Criação de um novo usuário
func CreateUser(user *models.User) error {
	// Verifica se o email já está cadastrado
	exists, err := VerifyUserCreated(user.Email)
	if err != nil {
		return err // Retorna erro em caso de problemas no banco
	}

	if exists {
		return errors.New("email already in use") // Retorna erro se o email já existe
	}

	// Cria o usuário no banco de dados
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}