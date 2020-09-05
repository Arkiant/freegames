package freegames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractCommand(t *testing.T) {
	useCase := []struct {
		name    string
		content string
		errors  bool
		args    int
		command string
	}{
		{
			name:    "correct values",
			content: "!freegames",
			errors:  false,
			args:    0,
			command: "freegames",
		},
		{
			name:    "correct values with args",
			content: "!join #test",
			errors:  false,
			args:    1,
			command: "join",
		},
		{
			name:    "wrong command",
			content: "/join #test",
			errors:  true,
			args:    0,
			command: "",
		},
	}

	for _, tc := range useCase {
		t.Run(tc.name, func(t *testing.T) {
			command, args, err := ExtractCommand(tc.content)
			if tc.errors {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Len(t, args, tc.args)
			assert.Equal(t, tc.command, command)
		})
	}
}
