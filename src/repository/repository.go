package repository

import (
	"github.com/WeslleyRibeiro-1999/login-go/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	SingUp(user *models.User) (*models.UserResponse, error)
	SignIn(login *models.UserLogin) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

var _ Repository = (*repository)(nil)

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) SingUp(user *models.User) (*models.UserResponse, error) {
	passwordHash, err := CriptografarSenha(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordHash)

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	userResponse := &models.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return userResponse, nil
}

func CriptografarSenha(password string) (string, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(newPassword), nil
}

func (r *repository) SignIn(login *models.UserLogin) (*models.User, error) {
	var user models.User

	password, err := CriptografarSenha(login.Password)
	if err != nil {
		return nil, err
	}

	if err := r.db.Take(&models.User{Email: login.Email, Password: string(password)}).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
