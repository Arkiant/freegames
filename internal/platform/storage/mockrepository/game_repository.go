// Code generated by mockery v2.9.4. DO NOT EDIT.

package mockrepository

import (
	freegames "github.com/arkiant/freegames/internal"
	mock "github.com/stretchr/testify/mock"
)

// GameRepository is an autogenerated mock type for the GameRepository type
type GameRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: game
func (_m *GameRepository) Delete(game freegames.Game) error {
	ret := _m.Called(game)

	var r0 error
	if rf, ok := ret.Get(0).(func(freegames.Game) error); ok {
		r0 = rf(game)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: game
func (_m *GameRepository) Exists(game freegames.Game) bool {
	ret := _m.Called(game)

	var r0 bool
	if rf, ok := ret.Get(0).(func(freegames.Game) bool); ok {
		r0 = rf(game)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetGames provides a mock function with given fields: platform
func (_m *GameRepository) GetGames(platform freegames.Platform) (freegames.FreeGames, error) {
	ret := _m.Called(platform)

	var r0 freegames.FreeGames
	if rf, ok := ret.Get(0).(func(freegames.Platform) freegames.FreeGames); ok {
		r0 = rf(platform)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(freegames.FreeGames)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(freegames.Platform) error); ok {
		r1 = rf(platform)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: game
func (_m *GameRepository) Store(game freegames.Game) error {
	ret := _m.Called(game)

	var r0 error
	if rf, ok := ret.Get(0).(func(freegames.Game) error); ok {
		r0 = rf(game)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
