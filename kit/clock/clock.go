package clock

import (
	"time"
)

type Clock interface {
	Now() time.Time
	After(d time.Duration) <-chan time.Time
}

func NewClock(t *time.Time) Clock {
	return ClockStub{
		Date: t,
	}
}

func NewParseTime(t string) time.Time {
	r, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return time.Time{}
	}

	return r
}

type ClockStub struct {
	Date *time.Time
}

func (c ClockStub) Now() time.Time {
	if c.Date != nil {
		return *c.Date
	}

	return time.Now()
}

func (ClockStub) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Europe/Madrid".
func TimeIn(t time.Time, locationName string) (time.Time, error) {
	loc, err := time.LoadLocation(locationName)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
