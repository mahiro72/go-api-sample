package dto

import (
	"time"

	"github.com/mahiro72/go-api-sample/domain/model"
)

type User struct {
	ID        string
	Name      string
	Password  string
	CreatedAt time.Time
}

func NewUser(m model.User) User {
	return User{
		ID:        m.ID.String(),
		Name:      m.Name,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
	}
}

func (dto User) ToModel() model.User {
	return model.User{
		ID:        model.UserID(dto.ID),
		Name:      dto.Name,
		Password:  dto.Password,
		CreatedAt: dto.CreatedAt,
	}
}
