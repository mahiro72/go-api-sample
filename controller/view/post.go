package view

import (
	"time"

	"github.com/mahiro72/go-api-sample/domain/model"
)

type post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPost(m model.Post) post {
	return post{
		ID:        m.ID.String(),
		Title:     m.Title,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
	}
}

func NewPostList(ms []model.Post) []post {
	r := []post{}

	for _, m := range ms {
		r = append(r, NewPost(m))
	}
	return r
}
