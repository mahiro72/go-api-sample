package repository

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
)

var ErrNoUserFound = errors.New("no user found")

type User interface {
	Create(ctx context.Context, user model.User) error
	Find(ctx context.Context, userID model.UserID) (model.User, error)
	FindByName(ctx context.Context, name string) (model.User, error)
	FindByNameAndPassword(ctx context.Context, name, password string) (model.User, error)
}
