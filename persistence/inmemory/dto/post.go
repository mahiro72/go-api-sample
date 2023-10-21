package dto

import (
	"time"

	"github.com/mahiro72/go-api-sample/domain/model"
)

type Post struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
}

func NewPost(m model.Post, userID model.UserID) Post {
	return Post{
		ID:        m.ID.String(),
		UserID:    userID.String(),
		Title:     m.Title,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
	}
}

func (dto Post) ToModel() model.Post {
	return model.Post{
		ID:        model.PostID(dto.ID),
		Title:     dto.Title,
		Content:   dto.Content,
		CreatedAt: dto.CreatedAt,
	}
}
