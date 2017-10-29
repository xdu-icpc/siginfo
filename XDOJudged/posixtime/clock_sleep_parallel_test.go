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
	"runtime"
	"testing"
	"time"

	"linux.xidian.edu.cn/git/XDU_ACM_ICPC/XDOJ-next/XDOJudged/posixtime"
)

func routine(ch chan<- error) {
	ch <- posixtime.CLOCK_MONOTONIC.Sleep(time.Second)
}

// Create many goroutines (10 times of GOMAXPROC) and confirm they can
// sleep simutaniously.
func TestSleepParallel(t *testing.T) {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	t0 := time.Now()

	ch := make(chan error)
	for i := 0; i <= 10*numCPU; i++ {
		go routine(ch)
	}

	for i := 0; i <= 10*numCPU; i++ {
		err := <-ch
		if err != nil {
			t.Errorf("Can not Sleep on CLOCK_MONOTOINC: %v", err)
		}
	}

	d := time.Now().Sub(t0)
	if d > time.Second*3/2 {
		t.Errorf("Go routines failed to sleep simutaniously.")
	}
}
