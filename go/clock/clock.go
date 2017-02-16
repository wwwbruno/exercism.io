package clock

import "fmt"

const (
	testVersion = 4
	minutesSize = 60
	hourSize    = 24
)

type Clock int

func New(hour, minute int) Clock {
	time := minute + (hour * minutesSize)
	for time < 0 {
		time = time + (hourSize * minutesSize)
	}
	time = time % (hourSize * minutesSize)

	return Clock(time)
}

func (clock Clock) String() string {
	hours, calculatedMinutes := (clock / minutesSize), (clock % minutesSize)
	calculatedHours := hours % hourSize

	return fmt.Sprintf("%02d:%02d", calculatedHours, calculatedMinutes)
}

func (clock Clock) Add(minutes int) Clock {
	time := clock + Clock(minutes)
	for time < 0 {
		time = time + (hourSize * minutesSize)
	}

	return Clock(time)
}
