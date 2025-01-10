// src/models/teachers.go
package models

import "time"

type Teachers struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	CPF         string     `gorm:"not null" json:"cpf"`
	BirthDate   string     `gorm:"not null" json:"birth_date"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
}