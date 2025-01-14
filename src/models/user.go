package models

type User struct {
    ID       uint   `gorm:"primary_key"     json:"id"`
    Username string `gorm:"unique;not null" json:"username" validate:"required,min=3,max=20"`
    Email    string `gorm:"unique;not null" json:"email" validate:"required,email"`
    Password string `gorm:"not null"        json:"password" validate:"required,min=8"`
}

type UserCreate struct {
    Username string `gorm:"unique;not null" json:"username" validate:"required,min=3,max=20" example:"usuario123"`
    Email    string `gorm:"unique;not null" json:"email" validate:"required,email" example:"usuario@example.com"`   
    Password string `gorm:"not null"        json:"password" validate:"required,min=8" example:"senhaSegura123"`     
}

type UserResponse struct {
    ID       uint   `gorm:"primary_key"     json:"id" example:"123"`       
    Username string `gorm:"unique;not null" json:"username" example:"usuario123"` 
    Email    string `gorm:"unique;not null" json:"email" example:"usuario@example.com"`    
    Password string `gorm:"not null"        json:"password" example:"senhaSegura123"` 
}

type UserUpdateRequest struct {
    Username string `gorm:"unique;not null" json:"username" validate:"min=3,max=20" example:"usuario123"` 
    Email    string `gorm:"unique;not null" json:"email" validate:"email" example:"usuario@example.com"`
    Password string `gorm:"not null"        json:"password" validate:"min=8" example:"senhaSegura123"`
}