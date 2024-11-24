package domain

import (
	"time"

	"github.com/wazwki/WearStore/user-service/api/proto/userpb"
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func EntityToDTO(user *User) *userpb.User {
	return &userpb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.GoString(),
		UpdatedAt: user.UpdatedAt.GoString(),
	}
}
