package models


import (
	"errors"
	"golang.org/x/crypto/bcrypt"
    "mvc/src/repositories"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
    Email  string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
}




func CreatedUser(user *User) (bool, error) {
	// Chama a função da camada repository
	err := repositories.CreateUser(user)
	if err != nil {
		if err.Error() == "email already in use" {
			return false, err 
		}
		return false, err 
	}

	return true, nil 
}

// Gerar hash da senha
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
