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

type SearchParams struct {
	Name    *string
	DOB     *string
	Address *string
	Country *string
	Gender  *string
	UserID  *string
	From    int64
	Size    int64
}

// SearchResults defines the collection of tasks that were found.
type SearchResults struct {
	Tasks []Task
	Total int64
}

// IsZero determines whether the search arguments have values or not.
func (a SearchParams) IsZero() bool {
	return a.Name == nil &&
		a.Country == nil &&
		a.DOB == nil &&
		a.Address == nil &&
		a.UserID == nil &&
		a.Gender == nil
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}
