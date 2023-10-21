package service

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/repository"
)

type userNameValidator struct {
	repoUser repository.User
}

func NewUserNameValidator(repoUser repository.User) *userNameValidator {
	return &userNameValidator{
		repoUser: repoUser,
	}
}

func (svc *userNameValidator) CheckAlreadyUsed(ctx context.Context, name string) (bool, error) {
	_, err := svc.repoUser.FindByName(ctx, name)
	if errors.Is(err, repository.ErrNoUserFound) {
		return false, nil
	}
	return true, err
}
