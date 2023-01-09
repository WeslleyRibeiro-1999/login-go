package main

import (
	"log"

	"github.com/WeslleyRibeiro-1999/login-go/database"
	"github.com/WeslleyRibeiro-1999/login-go/src/api"
	repository "github.com/WeslleyRibeiro-1999/login-go/src/repository"
	"github.com/labstack/echo/v4"
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

	repository.NewRepository(database)

	e := echo.New()
	e.POST("/singup", api.CreateUser)
	e.Logger.Fatal(e.Start(":8083"))
}
