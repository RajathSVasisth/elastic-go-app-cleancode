package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Index = "tasks"
)

type Task struct {
	ID      primitive.ObjectID `json:"-"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	DOB     string             `json:"dob"`
	Country string             `json:"country"`
	Gender  string             `json:"gender"`
	UserID  primitive.ObjectID `json:"-"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}
