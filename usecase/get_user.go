package usecase

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
)

type getUser struct {
	repoUser repository.User
}

func NewGetUser(repoUser repository.User) *getUser {
	return &getUser{
		repoUser: repoUser,
	}
}

func (uc *getUser) Exec(ctx context.Context, userIDString string) (model.User, error) {
	userID, err := model.ParseUserID(userIDString)
	if err != nil {
		if errors.Is(err, model.ErrUserIDValidation) {
			return model.User{}, ErrIDValidation
		}
		return model.User{}, err
	}

	user, err := uc.repoUser.Find(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNoUserFound) {
			return model.User{}, ErrNotExistsData
		}
		return model.User{}, err
	}
	return user, nil
}
