package model

import (
	"time"

	"github.com/google/uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.NewString())
}

func (u UserID) String() string {
	return string(u)
}

func (u UserID) IsSame(other UserID) bool {
	return u.String() == other.String()
}

func ParseUserID(userIDString string) (UserID,error) {
	userID,err := uuid.Parse(userIDString)
	if err != nil {
		return "",ErrUserIDValidation
	}
	return UserID(userID.String()),nil
}

type User struct {
	ID        UserID
	Name      string
	CreatedAt time.Time
}

func NewUser(name string) (User, error) {
	if len(name) > 100 {
		return User{}, ErrUserFieldsValidation
	}

	return User{
		ID:        NewUserID(),
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}
