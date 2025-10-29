package main

import (
	"ToDoList/internal/adapters/primary/api"
	"ToDoList/internal/adapters/secondary/postgresql"
	"ToDoList/internal/domain/service"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetPostgreEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить .env файл, используются системные переменные")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)
}

func main() {
	taskRepository, err := postgresql.NewPostgresqlRepository(GetPostgreEnv())
	if err != nil {
		panic(err)
	}
	taskService := service.NewTaskService(taskRepository)
	primaryAdapter := api.NewRouter(taskService)
	primaryAdapter.Start()
}
