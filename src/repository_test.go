package src

import (
	"testing"

	"github.com/WeslleyRibeiro-1999/login-go/database"
	"github.com/WeslleyRibeiro-1999/login-go/models"
	"github.com/stretchr/testify/assert"
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
			assert.Nil(t, err, "tem que ser vazio")

			assert.Equal(t, tt.want.ID, user.ID, "devem ser iguais")
			assert.Equal(t, tt.want.Email, user.Email, "devem ser iguais")
			assert.Equal(t, tt.want.Password, user.Password, "devem ser iguais")

			repository.db.Delete(models.User{ID: tt.args.user.ID})
		})
	}
}
