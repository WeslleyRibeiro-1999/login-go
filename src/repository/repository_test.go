package repository

import (
	"testing"

	"github.com/WeslleyRibeiro-1999/login-go/database"
	"github.com/WeslleyRibeiro-1999/login-go/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCriptografarSenha(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "as senhas devem ser iguais",
			arg:  "abc123",
			want: "abc123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := CriptografarSenha(tt.arg)
			assert.NoError(t, err, "error ao encriptar a senha")
			err = bcrypt.CompareHashAndPassword([]byte(password), []byte(tt.want))
			assert.NoError(t, err, "senha incorreta")
		})
	}
}

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
		want *models.UserResponse
	}{
		{
			name: "devem ser iguais",
			args: args{
				user: &models.User{
					FirstName: "Weslley",
					LastName:  "Jose",
					Email:     "teste@teste.com",
					Password:  "teste123",
				},
			},
			want: &models.UserResponse{
				FirstName: "Weslley",
				LastName:  "Jose",
				Email:     "teste@teste.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := repository.SingUp(tt.args.user)
			assert.NoError(t, err, "tem que ser vazio")

			assert.Equal(t, tt.want.Email, user.Email, "devem ser iguais")
			assert.Equal(t, tt.want.FirstName, user.FirstName)
			assert.Equal(t, tt.want.LastName, user.LastName)
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

			assert.Equal(t, tt.want.Email, user.Email, "emails devem ser iguais")
			passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tt.want.Password))
			assert.NoError(t, passwordErr, "senha nao retornada conforme esperada")
		})
	}

}

func TestDropTable(t *testing.T) {
	db, err := database.NewDatabase()
	assert.NoError(t, err, "sem erro ao inicializar o banco de dados")
	db.Migrator().DropTable(&models.User{})
}
