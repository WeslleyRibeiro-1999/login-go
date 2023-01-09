package repository

import (
	"testing"

	"github.com/WeslleyRibeiro-1999/login-go/database"
	"github.com/WeslleyRibeiro-1999/login-go/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestRepository_SingUp(t *testing.T) {
	db, err := database.NewDatabase()
	assert.NoError(t, err, "sem erro ao inicializar o banco de dados")
	repository := NewRepository(db)

	type args struct {
		user *models.User
	}
	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "devem ser iguais",
			args: args{
				user: &models.User{
					ID:        20,
					FirstName: "Weslley",
					LastName:  "Jose",
					Email:     "teste@teste.com",
					Password:  "teste123",
				},
			},
			want: &models.User{
				ID:        20,
				FirstName: "Weslley",
				LastName:  "Jose",
				Email:     "teste@teste.com",
				Password:  "teste123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := repository.SingUp(tt.args.user)
			assert.NoError(t, err, "tem que ser vazio")

			assert.Equal(t, tt.want.ID, user.ID, "devem ser iguais")
			assert.Equal(t, tt.want.Email, user.Email, "devem ser iguais")
			passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tt.want.Password))
			assert.NoError(t, passwordErr, "senha nao retornada conforme esperada")
		})
	}
}

func TestSignIn(t *testing.T) {
	db, err := database.NewDatabase()
	assert.NoError(t, err, "sem erro ao inicializar o banco de dados")
	repository := NewRepository(db)

	tests := []struct {
		name string
		args *models.UserLogin
		want *models.User
	}{
		{
			name: "deve buscar usuario pelo email e senha",
			args: &models.UserLogin{
				Email:    "teste@teste.com",
				Password: "teste123",
			},
			want: &models.User{
				ID:        20,
				FirstName: "Weslley",
				LastName:  "Jose",
				Email:     "teste@teste.com",
				Password:  "teste123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := repository.SignIn(tt.args)
			assert.NoError(t, err, "erro ao buscar usuario")

			assert.Equal(t, tt.want.ID, user.ID, "IDs devem ser iguais")
			assert.Equal(t, tt.want.Email, user.Email, "emails devem ser iguais")
			passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tt.want.Password))
			assert.NoError(t, passwordErr, "senha nao retornada conforme esperada")
		})
	}

}

func TestDropTable(t *testing.T) {
	db, err := database.NewDatabase()
	assert.NoError(t, err, "sem erro ao inicializar o banco de dados")
	repository := NewRepository(db)
	repository.db.Migrator().DropTable(&models.User{})
}
