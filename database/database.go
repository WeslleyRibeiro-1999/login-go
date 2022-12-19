package database

import (
	"github.com/WeslleyRibeiro-1999/login-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database.AutoMigrate(&models.User{})

	return database, nil
}
