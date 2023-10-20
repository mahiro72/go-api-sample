package persistence

import "github.com/mahiro72/go-api-sample/domain/model"

func ResetDatabase() error {
	database = Database{
		UserList: []model.User{},
		TaskList: []model.Task{},
	}
	return nil
}
