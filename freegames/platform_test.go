package freegames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPlatform(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []Platform
		Output int
	}{
		{
			Name:   "Add a single platform",
			Input:  []Platform{NewNoOpPlatform()},
			Output: 1,
		},
		{
			Name:   "Add three platforms",
			Input:  []Platform{NewNoOpPlatform(), NewNoOpPlatform(), NewNoOpPlatform()},
			Output: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			db, err := NewNoOpRepository()
			assert.NoError(t, err)

			fg := NewFreeGames(&db)
			for _, v := range tc.Input {
				fg.AddPlatform(v)
			}

			assert.Len(t, fg.platforms, tc.Output)
		})
	}
}
