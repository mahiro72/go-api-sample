package model

import (
	"time"

	"github.com/google/uuid"
)

type PostID string

func NewPostID() PostID {
	return PostID(uuid.NewString())
}

func (p PostID) String() string {
	return string(p)
}

func (p PostID) IsSame(other PostID) bool {
	return p.String() == other.String()
}

type Post struct {
	//投稿を一意に識別するID
	ID PostID
	//投稿のタイトル
	Title string
	//投稿にの内容
	Content string
	//投稿の作成日
	CreatedAt time.Time
}

func NewPost(title, content string) (Post, error) {
	if len(title) > 200 {
		return Post{}, ErrUserFieldsValidation
	}
	if len(content) > 600 {
		return Post{}, ErrUserFieldsValidation
	}

	return Post{
		ID:        NewPostID(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
