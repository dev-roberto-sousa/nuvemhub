package config

import (
	"fmt"
	"log"
	"os"

	"github.com/dev-roberto-sousa/nuvemhub/internal/api/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado.")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Erro ao executar AutoMigrate:", err)
	}

	DB = db
	log.Println("Conexão com o banco de dados realizada com sucesso!")
}
