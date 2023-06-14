// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mongo "github.com/RajathSVasisth/elasticApp/mongo"
	mock "github.com/stretchr/testify/mock"

	mongo_drivermongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is an autogenerated mock type for the Collection type
type Collection struct {
	mock.Mock
}

// Aggregate provides a mock function with given fields: _a0, _a1
func (_m *Collection) Aggregate(_a0 context.Context, _a1 interface{}) (mongo.Cursor, error) {
	ret := _m.Called(_a0, _a1)

	var r0 mongo.Cursor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (mongo.Cursor, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) mongo.Cursor); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Cursor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDocuments provides a mock function with given fields: _a0, _a1, _a2
func (_m *Collection) CountDocuments(_a0 context.Context, _a1 interface{}, _a2 ...*options.CountOptions) (int64, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) (int64, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) int64); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.CountOptions) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOne provides a mock function with given fields: _a0, _a1
func (_m *Collection) DeleteOne(_a0 context.Context, _a1 interface{}) (int64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (int64, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: _a0, _a1, _a2
func (_m *Collection) Find(_a0 context.Context, _a1 interface{}, _a2 ...*options.FindOptions) (mongo.Cursor, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongo.Cursor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOptions) (mongo.Cursor, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOptions) mongo.Cursor); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Cursor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.FindOptions) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: _a0, _a1
func (_m *Collection) FindOne(_a0 context.Context, _a1 interface{}) mongo.SingleResult {
	ret := _m.Called(_a0, _a1)

	var r0 mongo.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) mongo.SingleResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.SingleResult)
		}
	}

	return r0
}

// InsertMany provides a mock function with given fields: _a0, _a1
func (_m *Collection) InsertMany(_a0 context.Context, _a1 []interface{}) ([]interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}) ([]interface{}, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}) []interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOne provides a mock function with given fields: _a0, _a1
func (_m *Collection) InsertOne(_a0 context.Context, _a1 interface{}) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (interface{}, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMany provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Collection) UpdateMany(_a0 context.Context, _a1 interface{}, _a2 interface{}, _a3 ...*options.UpdateOptions) (*mongo_drivermongo.UpdateResult, error) {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo_drivermongo.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo_drivermongo.UpdateResult, error)); ok {
		return rf(_a0, _a1, _a2, _a3...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo_drivermongo.UpdateResult); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo_drivermongo.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOne provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Collection) UpdateOne(_a0 context.Context, _a1 interface{}, _a2 interface{}, _a3 ...*options.UpdateOptions) (*mongo_drivermongo.UpdateResult, error) {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo_drivermongo.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo_drivermongo.UpdateResult, error)); ok {
		return rf(_a0, _a1, _a2, _a3...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo_drivermongo.UpdateResult); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo_drivermongo.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCollection interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollection creates a new instance of Collection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollection(t mockConstructorTestingTNewCollection) *Collection {
	mock := &Collection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
