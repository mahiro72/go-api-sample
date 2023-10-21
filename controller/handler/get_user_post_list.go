package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mahiro72/go-api-sample/controller/response"
	"github.com/mahiro72/go-api-sample/impl/repository"
	"github.com/mahiro72/go-api-sample/persistence/inmemory"
	"github.com/mahiro72/go-api-sample/usecase"

	"github.com/go-chi/chi/v5"
)

func GetUserPostList(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "userID")

	db := inmemory.NewDatabase()
	uc := usecase.NewGetUserPostList(repository.NewUser(db), repository.NewPost(db))
	postList, err := uc.Exec(r.Context(), userIDString)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrIDValidation):
			w.WriteHeader(http.StatusBadRequest)
			return
		case errors.Is(err, usecase.ErrNoExistsData):
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	res := response.NewPostList(postList)
	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
