package freegames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddClient(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []Client
		Output int
	}{
		{
			Name:   "Add a single client",
			Input:  []Client{NewNoOpClient()},
			Output: 1,
		},
		{
			Name:   "Add three clients",
			Input:  []Client{NewNoOpClient(), NewNoOpClient(), NewNoOpClient()},
			Output: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			db, err := NewNoOpRepository()
			assert.NoError(t, err)

			fg := NewFreeGames(&db)
			for _, v := range tc.Input {
				fg.AddClient(v)
			}

			assert.Len(t, fg.clients, tc.Output)
		})
	}
}
