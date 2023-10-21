package persistence

import "github.com/mahiro72/go-api-sample/domain/model"

type Database struct {
	UserList []model.User
	TaskList []model.Task
}

var database Database

func init() {
	database = Database{
		UserList: []model.User{},
		TaskList: []model.Task{},
	}
}

func NewDatabase() *Database {
	return &database
}
