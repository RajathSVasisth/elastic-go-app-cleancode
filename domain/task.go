package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Index = "tasks"
)

type Task struct {
	ID      primitive.ObjectID `json:"-" form:"-"`
	Name    string             `json:"name" form:"name"`
	Address string             `json:"address" form:"address"`
	DOB     string             `json:"dob" form:"dob"`
	Country string             `json:"country" form:"country"`
	Gender  string             `json:"gender" form:"gender"`
	UserID  primitive.ObjectID `json:"userid" form:"-"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}
