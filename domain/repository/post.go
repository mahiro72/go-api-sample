package repository

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
)

var ErrNoPostFound = errors.New("no post found")

type Post interface {
	Create(ctx context.Context, post model.Post, userID model.UserID) error
	Find(ctx context.Context, postID model.PostID) (model.Post, error)
	FindAllByUserID(ctx context.Context, userID model.UserID) ([]model.Post, error)
}
