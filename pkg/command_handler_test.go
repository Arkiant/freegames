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
		args    string
		command string
	}{
		{
			name:    "correct values",
			content: "!freegames",
			errors:  false,
			args:    "",
			command: "freegames",
		},
		{
			name:    "correct values with args",
			content: "!join #test",
			errors:  false,
			args:    "#test",
			command: "join",
		},
		{
			name:    "wrong command",
			content: "/join #test",
			errors:  true,
			args:    "",
			command: "",
		},
	}

	for _, tc := range useCase {
		t.Run(tc.name, func(t *testing.T) {
			command, args, err := ParseCommand(tc.content)
			if tc.errors {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.args, args)
			assert.Equal(t, tc.command, command)
		})
	}
}
