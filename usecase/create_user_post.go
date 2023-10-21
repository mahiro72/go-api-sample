package usecase

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
	"github.com/mahiro72/go-api-sample/domain/service"
)

type createUserPost struct {
	repoPost                   repository.Post
	svcUserRegistrationChecker service.UserRegistrationChecker
}

func NewCreateUserPost(repoPost repository.Post, svcUserRegistrationChecker service.UserRegistrationChecker) *createUserPost {
	return &createUserPost{
		repoPost:                   repoPost,
		svcUserRegistrationChecker: svcUserRegistrationChecker,
	}
}

func (uc *createUserPost) Exec(ctx context.Context, title, content, name, password string) (model.Post, error) {
	// userが登録されているか nameとpasswordを使って確認
	userID, err := uc.svcUserRegistrationChecker.CheckUserRegistration(ctx, name, password)
	if err != nil {
		if errors.Is(err, service.ErrNoUserFound) {
			return model.Post{}, ErrNoExistsData
		}
		return model.Post{}, err
	}

	post, err := model.NewPost(title, content)
	if err != nil {
		if errors.Is(err, model.ErrPostFieldsValidation) {
			return model.Post{}, ErrFieldsValidation
		}
		return model.Post{}, err
	}

	err = uc.repoPost.Create(ctx, post, userID)
	if err != nil {
		return model.Post{}, err
	}
	return post, nil
}
