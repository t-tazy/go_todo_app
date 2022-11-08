package clock

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (r RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocer struct{}

func (fc FixedClocer) Now() time.Time {
	return time.Date(2022, 5, 10, 12, 34, 56, 0, time.UTC)
}
