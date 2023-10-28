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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var j requestCreateUser
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := inmemory.NewDatabase()
	repoUser := repository.NewUser(db)
	uc := usecase.NewCreateUser(repoUser, service.NewUserNameValidator(repoUser))
	user, err := uc.Exec(r.Context(), j.Name, j.Password)
	if err != nil {
		if errors.Is(err, usecase.ErrFieldsValidation) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := view.NewUser(user)
	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

type requestCreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
