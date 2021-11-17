// Code generated by mockery v2.9.4. DO NOT EDIT.

package querymocks

import (
	context "context"

	command "github.com/arkiant/freegames/kit/query"

	mock "github.com/stretchr/testify/mock"
)

// Bus is an autogenerated mock type for the Bus type
type Bus struct {
	mock.Mock
}

// Dispatch provides a mock function with given fields: _a0, _a1
func (_m *Bus) Dispatch(_a0 context.Context, _a1 command.Query) []byte {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, command.Query) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Register provides a mock function with given fields: _a0, _a1
func (_m *Bus) Register(_a0 command.Type, _a1 command.Handler) {
	_m.Called(_a0, _a1)
}
