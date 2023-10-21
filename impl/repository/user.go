package repository

import (
	"context"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
	"github.com/mahiro72/go-api-sample/persistence/inmemory"
	"github.com/mahiro72/go-api-sample/persistence/inmemory/dto"
)

type user struct {
	db *inmemory.Database
}

func NewUser(db *inmemory.Database) repository.User {
	return &user{
		db: db,
	}
}

func (repo *user) Create(ctx context.Context, user model.User) error {
	repo.db.UserList = append(repo.db.UserList, dto.NewUser(user))
	return nil
}

func (repo *user) Find(ctx context.Context, userID model.UserID) (model.User, error) {
	for _, u := range repo.db.UserList {
		user := u.ToModel()
		if user.ID.IsSame(userID) {
			return user, nil
		}
	}
	return model.User{}, repository.ErrNoUserFound
}

func (repo *user) FindByName(ctx context.Context, name string) (model.User, error) {
	for _, u := range repo.db.UserList {
		user := u.ToModel()
		if name == user.Name {
			return user, nil
		}
	}
	return model.User{}, repository.ErrNoUserFound
}

func (repo *user) FindByNameAndPassword(ctx context.Context, name, password string) (model.User, error) {
	for _, u := range repo.db.UserList {
		user := u.ToModel()
		if name == user.Name && password == user.Password {
			return user, nil
		}
	}
	return model.User{}, repository.ErrNoUserFound
}
