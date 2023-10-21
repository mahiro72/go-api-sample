package usecase

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
	"github.com/mahiro72/go-api-sample/domain/service"
)

type createUser struct {
	repoUser             repository.User
	svcUserNameValidator service.UserNameValidator
}

func NewCreateUser(repoUser repository.User, svcUserNameValidator service.UserNameValidator) *createUser {
	return &createUser{
		repoUser:             repoUser,
		svcUserNameValidator: svcUserNameValidator,
	}
}

func (uc *createUser) Exec(ctx context.Context, name, password string) (model.User, error) {
	// nameが使われていないか確認
	used, err := uc.svcUserNameValidator.CheckAlreadyUsed(ctx, name)
	if err != nil {
		return model.User{}, err
	}
	if used {
		return model.User{}, ErrUserNameAlreadyUsed
	}

	user, err := model.NewUser(name, password)
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
