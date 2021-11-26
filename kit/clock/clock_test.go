package clock

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeIn(t *testing.T) {
	utcTime, _ := time.Parse("2006-01-02T15:04:05Z", "2021-08-06T23:00:00Z")

	var cases = []struct {
		name         string
		locationName string
		expected     string
		err          error
	}{
		{
			name:         "Error unknown timezone",
			locationName: "Unknown/Timezone",
			expected:     "",
			err:          errors.New("unknown time zone Unknown/Timezone"),
		},
		{
			name:         "With location name Europe/Madrid",
			locationName: "Europe/Madrid",
			expected:     "2021-08-07T01:00:00Z",
			err:          nil,
		},
		{
			name:         "With location name America/Argentina/Cordoba",
			locationName: "America/Argentina/Cordoba",
			expected:     "2021-08-06T20:00:00Z",
			err:          nil,
		},
		{
			name:         "With location name America/Argentina/Cordoba",
			locationName: "America/Argentina/Cordoba",
			expected:     "2021-08-06T20:00:00Z",
			err:          nil,
		},
		{
			name:         "With location name Pacific/Honolulu",
			locationName: "Pacific/Honolulu",
			expected:     "2021-08-06T13:00:00Z",
			err:          nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			locationTimezone, err := TimeIn(utcTime, c.locationName)
			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, c.err.Error())
			}
			if err == nil {
				locationTimezoneToString := locationTimezone.Format("2006-01-02T15:04:05Z")
				assert.Equal(t, c.expected, locationTimezoneToString)
			}
		})
	}

}
