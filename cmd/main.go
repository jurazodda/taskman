package main

import (
	"log"
	"taskman/config"
	"taskman/internal/handler"
	"taskman/internal/repository"
	"taskman/internal/service"
)

func main() {
	dsn, err := config.InitConfig()
	if err != nil {
		log.Fatal("error initializing configuration", err)
	}

	db, err := repository.DBConnection(dsn)
	if err != nil {
		log.Fatal("failed to initialize DBConnection", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewTaskService(repos)
	handlers := handler.NewHandler(services)

	log.Println("Server is listening...")
	if err := handler.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("failed running server", err)
	}
}
