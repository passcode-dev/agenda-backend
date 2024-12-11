package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

func LoadEnv() {
    // Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }
}

func ConnectDB() (*sql.DB, error) {
    LoadEnv() // Certifica-se de carregar as variáveis de ambiente

    // Recupera os dados do .env
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Formata a string de conexão
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("erro ao abrir a conexão com o banco: %v", err)
    }

    // Testa a conexão
    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
    }

    fmt.Println("Conectado ao banco de dados:", dbName)
    return db, nil
}
