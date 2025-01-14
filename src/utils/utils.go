package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword gera um hash bcrypt a partir de uma senha e retorna o hash ou um erro
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}
