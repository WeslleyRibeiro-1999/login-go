package repository

import (
	"github.com/WeslleyRibeiro-1999/login-go/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SingUp(user *models.User) (*models.User, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(newPassword)

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) SignIn(login *models.UserLogin) (*models.User, error) {
	var user models.User

	password, err := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if err := r.db.Take(&models.User{Email: login.Email, Password: string(password)}).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
