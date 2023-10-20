package usecase

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
)

type createUser struct {
	repoUser repository.User
}

func NewCreateUser(repoUser repository.User) *createUser {
	return &createUser{
		repoUser: repoUser,
	}
}

func (uc *createUser) Exec(ctx context.Context, name string) (model.User, error) {
	user, err := model.NewUser(name)
	if err != nil {
		if errors.Is(err, model.ErrUserFieldsValidation) {
			return model.User{}, ErrFieldsValidation
		}
		return model.User{}, err
	}

	err = uc.repoUser.Create(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
