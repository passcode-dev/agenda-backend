package models

import (
	"time"
	"gorm.io/gorm"
)

type Aluno struct {
	ID             uint           `gorm:"primary_key" json:"id"`
	Name           string         `gorm:"not null" json:"name"`
	DocumentRG     string         `gorm:"not null" json:"document_rg"`
	DocumentCPF    string         `gorm:"unique;not null" json:"document_cpf"`
	BirthDate      string         `gorm:"not null" json:"birth_date"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      *time.Time     `json:"deleted_at,omitempty"`
}