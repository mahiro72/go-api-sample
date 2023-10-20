package repository

import (
	"context"

	"github.com/mahiro72/go-api-sample/domain/model"
)

type User interface {
	Create(ctx context.Context, user model.User) error
	Find(ctx context.Context, userID model.UserID) (model.User, error)
}
