package service

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
	"github.com/mahiro72/go-api-sample/domain/service"
	"github.com/mahiro72/go-api-sample/utils"
)

type userRegistrationChecker struct {
	repoUser repository.User
}

func NewUserRegistrationChecker(repoUser repository.User) *userRegistrationChecker {
	return &userRegistrationChecker{
		repoUser: repoUser,
	}
}

func (svc *userRegistrationChecker) CheckUserRegistration(ctx context.Context, name, password string) (model.UserID, error) {
	hashPwd := utils.CreateHashFromString(name + password)

	user, err := svc.repoUser.FindByNameAndPassword(ctx, name, hashPwd)
	if err != nil {
		if errors.Is(err, repository.ErrNoUserFound) {
			return "", service.ErrNoUserFound
		}
		return "", err
	}
	return user.ID, nil
}
