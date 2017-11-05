// Test if ClockID.Sleep() works along with the Go runtime sched.
// Copyright (C) 2017  Laboratory of ACM/ICPC, Xidian University

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Author: Xi Ruoyao <ryxi@stu.xidian.edu.cn>

// +build linux

package posixtime_test

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func routine(d time.Duration, ch chan<- error) {
	ch <- sleepWell(d)
}

// Create many goroutines (10 times of GOMAXPROC) and confirm they can
// sleep simutaniously.
func TestSleepParallel(t *testing.T) {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	t0 := time.Now()

	ch := make(chan error)
	for i := 0; i <= 10*numCPU; i++ {
		go routine(sleepDuration, ch)
	}

	for i := 0; i <= 10*numCPU; i++ {
		err := <-ch
		if err != nil {
			t.Errorf("can not sleepWell: %v", err)
		}
	}

	d := time.Now().Sub(t0)
	if d > sleepDuration*3/2 {
		t.Errorf("goroutines failed to sleep simutaniously.")
	}
}

// Prove the runtime won't wake up a wrong goroutine even if we have only
// one OS thread.
func TestSleepDifferentTime(t *testing.T) {
	runtime.GOMAXPROCS(1)

	ch := make(chan error)
	keys := make([]int64, 10)
	chkey := make(chan int64, 10)

	for i := 0; i < 10; i++ {
		keys[i] = rand.Int63()
		key := keys[i]
		d := time.Second - time.Millisecond*time.Duration(i*100)
		go func() {
			localkey := key
			localkey = localkey ^ 0x12345678abcd
			routine(d, ch)
			localkey = localkey ^ 19260817
			chkey <- localkey
		}()
	}

	for i := 0; i < 10; i++ {
		err := <-ch
		if err != nil {
			t.Errorf("can not sleepWell: %v", err)
		}

		// We should receive the key of the goroutine slept less time
		// earlier.
		key := <-chkey
		key = key ^ 19260817 ^ 0x12345678abcd
		t.Logf("the %d-th key decoded is %d, saved key is %d",
			i, key, keys[9-i])
		if key != keys[9-i] {
			t.Errorf("wrong key!")
		}
	}
}
