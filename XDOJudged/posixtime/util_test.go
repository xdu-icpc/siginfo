package posixtime_test

import (
	"fmt"
	"time"
)

func checkDuration(infact time.Duration, should time.Duration) error {
	delta := infact.Nanoseconds() - should.Nanoseconds()
	if delta < 0 { // POSIX said it's impossible
		return fmt.Errorf("Slept too short: delta = %d ns", delta)
	}
	if delta > 50000000 { // max tolerance is 50ms
		return fmt.Errorf("Slept too long: delta = %d ns", delta)
	}

	return nil
}
