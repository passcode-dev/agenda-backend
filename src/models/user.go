package models

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
}

/*
func VerifyUser(username, password string) (*User, error){

    var user User

    database.DB.Where("username = ?", username).First(&user)
    if user.ID == 0 {
        log.Println("uário não encontrado:")
        return nil, errors.New("usuário não encontrado")
        
    }

    if(services.Authenticate(password,user.Password)){
        return &user, nil
    }

    log.Println("senha inválida")
    return nil, errors.New("senha inválida")

}*/
