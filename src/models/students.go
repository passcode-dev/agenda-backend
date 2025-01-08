// src/models/aluno.go
package models

import (
	"time"
)

type Students struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	RG          string     `gorm:"not null" json:"rg"`
	CPF         string     `gorm:"not null" json:"cpf"`
	BirthDate   string     `gorm:"not null" json:"birth_date"`
	PhoneNumber string     `gorm:"not null" json:"phone_number"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"` // Automático na criação
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"` // Nulo na criação, atualizado automaticamente
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`          // Nulo na criação, atualizado manualmente
}
