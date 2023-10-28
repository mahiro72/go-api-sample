package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mahiro72/go-api-sample/controller/view"
	"github.com/mahiro72/go-api-sample/impl/repository"
	"github.com/mahiro72/go-api-sample/impl/service"
	"github.com/mahiro72/go-api-sample/persistence/inmemory"
	"github.com/mahiro72/go-api-sample/usecase"
)

func CreateUserPost(w http.ResponseWriter, r *http.Request) {
	// 認証情報取得
	name, password, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var j requestCreateUserPost
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := inmemory.NewDatabase()
	svcUserRegistrationChecker := service.NewUserRegistrationChecker(repository.NewUser(db))
	uc := usecase.NewCreateUserPost(repository.NewPost(db), svcUserRegistrationChecker)
	post, err := uc.Exec(r.Context(), j.Title, j.Content, name, password)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrFieldsValidation):
			w.WriteHeader(http.StatusBadRequest)
			return
		case errors.Is(err, usecase.ErrNotExistsData):
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	res := view.NewPost(post)
	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

type requestCreateUserPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
