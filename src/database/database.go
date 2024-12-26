package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"agenda-backend/src/models"
	
	
)

var DB *gorm.DB

func InitDB() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, carregando variáveis de ambiente...")
	}

		// Acessando as variáveis de ambiente
		dbConnection := os.Getenv("DB_CONNECTION")
		jwtSecret := os.Getenv("JWT_SECRET")
	
		log.Printf("Conexão com o banco de dados: %s", dbConnection)
		log.Printf("JWT Secret: %s", jwtSecret)

	DB, err = gorm.Open("mysql", "teste:1234@tcp(localhost:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Falha ao conectar ao banco de dados:", err)
		panic("Erro ao conectar ao banco de dados")
	}

	if DB == nil {
		panic("Erro ao inicializar o banco de dados: ponteiro nulo")
	}

	if err = DB.AutoMigrate(
		&models.User{},
		/*&models.Student{},
		&models.Teacher{},
		&models.Class{},
		&models.Course{},
		&models.StudentHasClass{},
		&models.StudentHasCourse{},
		&models.TeacherHasCourse{},
		&models.TeacherHasClass{},*/
	).Error; err != nil {
		panic("Erro ao criar as tabelas do banco de dados: " + err.Error())
	}
}

var JWTSecret = os.Getenv("JWT_SECRET")
