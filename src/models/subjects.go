package models

import "time"

type Subjects struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}