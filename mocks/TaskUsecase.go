// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/RajathSVasisth/elastic-go-app-cleancode/domain"
	mock "github.com/stretchr/testify/mock"
)

// TaskUsecase is an autogenerated mock type for the TaskUsecase type
type TaskUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, task
func (_m *TaskUsecase) Create(c context.Context, task *domain.Task) error {
	ret := _m.Called(c, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) error); ok {
		r0 = rf(c, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, id
func (_m *TaskUsecase) Delete(c context.Context, id *string) error {
	ret := _m.Called(c, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) error); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchByUserID provides a mock function with given fields: c, userID, pagination
func (_m *TaskUsecase) FetchByUserID(c context.Context, userID string, pagination domain.Pagination) ([]domain.Task, error) {
	ret := _m.Called(c, userID, pagination)

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Pagination) ([]domain.Task, error)); ok {
		return rf(c, userID, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Pagination) []domain.Task); ok {
		r0 = rf(c, userID, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.Pagination) error); ok {
		r1 = rf(c, userID, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, task
func (_m *TaskUsecase) Update(c context.Context, task *domain.Task) error {
	ret := _m.Called(c, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) error); ok {
		r0 = rf(c, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTaskUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskUsecase creates a new instance of TaskUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskUsecase(t mockConstructorTestingTNewTaskUsecase) *TaskUsecase {
	mock := &TaskUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
