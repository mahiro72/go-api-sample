package repository

import (
	"context"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
	"github.com/mahiro72/go-api-sample/persistence"
)

type user struct {
	db *persistence.Database
}

func NewUser(db *persistence.Database) repository.User {
	return &user{
		db: db,
	}
}

func (repo *user) Create(ctx context.Context, user model.User) error {
	repo.db.UserList = append(repo.db.UserList, user)
	return nil
}

func (repo *user) Find(ctx context.Context, userID model.UserID) (model.User, error) {
	for _,u := range repo.db.UserList {
		if u.ID.IsSame(userID) {
			return u,nil
		}
	}
	return model.User{}, repository.ErrNoUserFound
}
