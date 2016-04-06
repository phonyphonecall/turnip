package turnip

import "time"

const (
	defaultName = "deafult"
)

type Timer struct {
	start time.Time

	times []time.Duration

	avg   float64
	count float64
}

// Start a timer
func (t *Timer) Start() {
	t.start = time.Now()
}

// Stop a timer, and update stats
func (t *Timer) Stop() {
	dur := time.Since(t.start)

	t.times = append(t.times, dur)
	t.count++

	// update average a1 = a0 + (t1 - a1)/n
	t.avg = t.avg + ((float64(dur.Nanoseconds()) - t.avg) / t.count)
}

// Get the average timer duration in Milliseconds
func (t *Timer) AverageMs() float64 {
	ms := t.avg / (float64(time.Millisecond))

	return ms
}

// Create a new timer
func NewTimer() *Timer {
	t := &Timer{
		times: make([]time.Duration, 0),
		count: 0,
	}

	return t
}
