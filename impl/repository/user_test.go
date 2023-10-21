package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mahiro72/go-api-sample/domain/model"
	"github.com/mahiro72/go-api-sample/domain/repository"
	repositoryImpl "github.com/mahiro72/go-api-sample/impl/repository"
	"github.com/mahiro72/go-api-sample/persistence"
)

func TestUser_Create(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		persistence.ResetDatabase()
		db := persistence.NewDatabase()

		repo := repositoryImpl.NewUser(db)
		err := repo.Create(context.Background(), model.User{
			ID:        "xxxx",
			Name:      "xxxx",
			CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestUser_Find(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		persistence.ResetDatabase()
		db := persistence.NewDatabase()
		// データの用意
		err := repositoryImpl.NewUser(db).Create(context.Background(), model.User{
			ID:        "xxxx",
			Name:      "xxxx",
			CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		})
		if err != nil {
			panic(err)
		}

		repo := repositoryImpl.NewUser(db)
		_, err = repo.Find(context.Background(), "xxxx")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("failure no user found", func(t *testing.T) {
		persistence.ResetDatabase()
		db := persistence.NewDatabase()
		// データの用意
		err := repositoryImpl.NewUser(db).Create(context.Background(), model.User{
			ID:        "xxxx",
			Name:      "xxxx",
			CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		})
		if err != nil {
			panic(err)
		}

		repo := repositoryImpl.NewUser(db)
		_, err = repo.Find(context.Background(), "yyyy")
		if !errors.Is(err, repository.ErrNoUserFound) {
			t.Fatal(err)
		}
	})
}
