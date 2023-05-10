package testutil

import "time"

// A utility to measure the running time of a program.
type Stopwatch struct {
	start time.Time
}

func NewStopwatch() Stopwatch {
	return Stopwatch{time.Now()}
}

// ElapsedTime method returns the elapsed time (in seconds) since the stopwatch
// was created.
func (s Stopwatch) ElapsedTime() float64 {
	now := time.Now()
	dt := now.Sub(s.start)
	return float64(dt.Milliseconds()) / 1000
}
