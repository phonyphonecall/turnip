package turnip

import (
	"testing"
	"time"
)

func TestStartStop(t *testing.T) {
	timer := NewTimer()

	for i := 0; i < 100; i++ {
		timer.Start()
		time.Sleep(20 * time.Millisecond)
		timer.Stop()

		// Allow +/- 5 ms
		dur := timer.AverageMs()
		if dur > 25 || dur < 15 {
			t.Fail()
		}
	}
}
