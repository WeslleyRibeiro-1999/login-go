package database

import (
	"database/sql"

	"github.com/WeslleyRibeiro-1999/login-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*sql.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database.AutoMigrate(&models.User{})

	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}

	return sqlDB, nil
}
