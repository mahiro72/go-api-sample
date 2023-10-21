package inmemory

import (
	"github.com/mahiro72/go-api-sample/persistence/inmemory/dto"
)

// test用のDatabaseリセット関数
func ResetDatabase() error {
	database = Database{
		UserList: []dto.User{},
		PostList: []dto.Post{},
	}
	return nil
}
