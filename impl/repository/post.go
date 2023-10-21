package repository

import (
	"context"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
	"github.com/mahiro72/go-api-sample/persistence/inmemory"
	"github.com/mahiro72/go-api-sample/persistence/inmemory/dto"
)

type post struct {
	db *inmemory.Database
}

func NewPost(db *inmemory.Database) *post {
	return &post{
		db: db,
	}
}

func (repo *post) Create(ctx context.Context, post model.Post, userID model.UserID) error {
	repo.db.PostList = append(repo.db.PostList, dto.NewPost(post, userID))
	return nil
}

func (repo *post) Find(ctx context.Context, postID model.PostID) (model.Post, error) {
	for _, p := range repo.db.PostList {
		post := p.ToModel()
		if post.ID.IsSame(postID) {
			return post, nil
		}
	}
	return model.Post{}, repository.ErrNoPostFound
}

func (repo *post) FindAllByUserID(ctx context.Context, userID model.UserID) ([]model.Post, error) {
	postList := []model.Post{}

	for _, p := range repo.db.PostList {
		if model.UserID(p.UserID).IsSame(userID) {
			postList = append(postList, p.ToModel())
		}
	}
	return postList, nil
}
