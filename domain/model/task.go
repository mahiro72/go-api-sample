package model

type Status string

const (
	NotStarted Status = "NotStarted"
	InProgress Status = "InProgress"
	Completed  Status = "Completed"
)

type TaskID string

type Task struct {
	ID     TaskID
	Name   string
	Status Status
}
