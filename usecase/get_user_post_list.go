package usecase

import (
	"context"
	"errors"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
)

type getUserPostList struct {
	repoUser repository.User
	repoPost repository.Post
}

func NewGetUserPostList(repoUser repository.User, repoPost repository.Post) *getUserPostList {
	return &getUserPostList{
		repoUser: repoUser,
		repoPost: repoPost,
	}
}

func (uc *getUserPostList) Exec(ctx context.Context, userIDString string) ([]model.Post, error) {
	userID, err := model.ParseUserID(userIDString)
	if err != nil {
		if errors.Is(err, model.ErrUserIDValidation) {
			return []model.Post{}, ErrIDValidation
		}
		return []model.Post{}, err
	}

	// userが存在するか確認
	_, err = uc.repoUser.Find(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNoUserFound) {
			return []model.Post{}, ErrNoExistsData
		}
		return []model.Post{}, err
	}

	post, err := uc.repoPost.FindAllByUserID(ctx, userID)
	if err != nil {
		return []model.Post{}, err
	}
	return post, nil
}
