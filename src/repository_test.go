package src

import (
	"testing"
	"time"

	"github.com/WeslleyRibeiro-1999/login-go/models"
	"github.com/stretchr/testify/assert"
)

func TestRepository_SingUp(t *testing.T) {

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
					ID:        1,
					FirstName: "Weslley",
					LastName:  "Jose",
					Email:     "teste@teste.com",
					Password:  "teste123",
					CreatedAt: time.Date(2010, time.October, 20, 10, 5, 30, 10, time.Local),
					UpdatedAt: time.Date(2010, time.October, 10, 10, 5, 30, 10, time.Local),
				},
			},
			want: &models.User{
				ID:        1,
				FirstName: "Weslley",
				LastName:  "Jose",
				Email:     "teste@teste.com",
				Password:  "teste123",
				CreatedAt: time.Date(2010, time.October, 20, 10, 5, 30, 10, time.Local),
				UpdatedAt: time.Date(2010, time.October, 10, 10, 5, 30, 10, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SingUp(tt.args.user)
			t.Logf("RETORNO: %v", got)
			assert.NotNil(t, err, "tem que ser vazio")
			assert.Equal(t, tt.want, got, "devem ser iguais")
		})
	}
}
