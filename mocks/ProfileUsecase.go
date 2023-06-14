// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/RajathSVasisth/elasticApp/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProfileUsecase is an autogenerated mock type for the ProfileUsecase type
type ProfileUsecase struct {
	mock.Mock
}

// GetProfileByID provides a mock function with given fields: c, userID
func (_m *ProfileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	ret := _m.Called(c, userID)

	var r0 *domain.Profile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Profile, error)); ok {
		return rf(c, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Profile); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Profile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProfileUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewProfileUsecase creates a new instance of ProfileUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProfileUsecase(t mockConstructorTestingTNewProfileUsecase) *ProfileUsecase {
	mock := &ProfileUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
