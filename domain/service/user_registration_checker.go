package service

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
)

var ErrNoUserFound = errors.New("no user found")

type UserRegistrationChecker interface {
	CheckUserRegistration(ctx context.Context, name, password string) (model.UserID, error)
}
