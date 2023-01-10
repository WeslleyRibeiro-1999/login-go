package main

import (
	"log"

	"github.com/WeslleyRibeiro-1999/login-go/database"
	"github.com/WeslleyRibeiro-1999/login-go/src/api"
	repository "github.com/WeslleyRibeiro-1999/login-go/src/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Erro ao inicializar database: %+v", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Erro ao inicializar sql: %+v", err)
	}
	defer sqlDB.Close()

	repo := repository.NewRepository(database)

	loginHandler := api.NewHandler(repo)

	app := fiber.New()
	app.Post("/signin", loginHandler.CreateUser)
	app.Listen(":3000")

}
