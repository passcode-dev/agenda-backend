package utils

import (
	"regexp"
	"errors"
)

func ValidateCPF(cpf string) error {
	// Remove todos os caracteres não numéricos
	numericOnly := regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")

	// Verifica se o CPF tem 11 dígitos
	if len(numericOnly) != 11 {
		return errors.New("CPF deve ter exatamente 11 dígitos")
	}

	// Valida os dígitos verificadores do CPF
	if !isValidCPF(numericOnly) {
		return errors.New("CPF inválido")
	}

	return nil
}

func isValidCPF(cpf string) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	firstVerifier := 11 - (sum % 11)
	if firstVerifier >= 10 {
		firstVerifier = 0
	}

	if firstVerifier != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	secondVerifier := 11 - (sum % 11)
	if secondVerifier >= 10 {
		secondVerifier = 0
	}

	return secondVerifier == int(cpf[10]-'0')
}
