package inmemory

import (
	"github.com/mahiro72/go-api-sample/persistence/inmemory/dto"
)

// FIXME: lockが必要
type Database struct {
	UserList []dto.User
	PostList []dto.Post
}

var database Database

func init() {
	database = Database{
		UserList: []dto.User{},
		PostList: []dto.Post{},
	}
}

func NewDatabase() *Database {
	return &database
}
