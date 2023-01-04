package src

import (
	"github.com/WeslleyRibeiro-1999/login-go/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SingUp(user *models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) SignIn(login *models.UserLogin) (*models.User, error) {
	var user models.User
	if err := r.db.Take(&models.User{Email: login.Email, Password: login.Password}).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
