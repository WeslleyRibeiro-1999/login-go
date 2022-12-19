package src

import (
	"github.com/WeslleyRibeiro-1999/login-go/database"
	"github.com/WeslleyRibeiro-1999/login-go/models"
)

func SingUp(user *models.User) (*models.User, error) {
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
