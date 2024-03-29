// Code generated by mockery v2.9.4. DO NOT EDIT.

package querymocks

import (
	context "context"

	query "github.com/arkiant/freegames/kit/cqrs/query"
	mock "github.com/stretchr/testify/mock"
)

// Bus is an autogenerated mock type for the Bus type
type Bus struct {
	mock.Mock
}

// Dispatch provides a mock function with given fields: _a0, _a1
func (_m *Bus) Dispatch(_a0 context.Context, _a1 query.Query) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, query.Query) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, query.Query) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0, _a1
func (_m *Bus) Register(_a0 query.Type, _a1 query.Handler) {
	_m.Called(_a0, _a1)
}
