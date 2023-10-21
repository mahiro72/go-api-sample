package router

import (
	"log"
	"net/http"

	"github.com/mahiro72/go-api-sample/controller/handler"

	"github.com/go-chi/chi/v5"
)

func Serve() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUser)
		r.Get("/{userID}", handler.GetUser)
	})

	r.Route("/users/{userID}/posts", func(r chi.Router) {
		r.Post("/", handler.CreateUserPost)
		r.Get("/list", handler.GetUserPostList)
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
