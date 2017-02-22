package gigasecond

import "time"

const (
	testVersion = 4
	gigaSecond  = 1000000000
)

func AddGigasecond(givenTime time.Time) time.Time {
	duration := time.Duration(gigaSecond) * time.Second
	return givenTime.Add(duration)
}
