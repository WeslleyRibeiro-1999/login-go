package main

import (
	"log"
	"net/http"

	"github.com/WeslleyRibeiro-1999/login-go/database"
	"github.com/WeslleyRibeiro-1999/login-go/src"
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

	src.NewRepository(database)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO WORLD!")
	})
	e.Logger.Fatal(e.Start(":8083"))
}
