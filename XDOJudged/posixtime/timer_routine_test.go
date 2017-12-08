// Whitebox test for timer_routine.go.
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

// white box test.

package posixtime

import (
	"testing"
	"time"

	"golang.org/x/sys/unix"
)

func TestTimerRoutine(t *testing.T) {
	t0 := time.Now()
	ts := unix.NsecToTimespec(1000000000)
	chstop := make(chan struct{})
	done := make(chan struct{})
	timerRoutine(CLOCK_MONOTONIC, ts, chstop, func() {
		d := time.Since(t0)
		t.Logf("d = %v", d)
		if d.Nanoseconds()-1000000000 > 10000000 {
			t.Fatalf("slept too long")
		}
		done <- struct{}{}
	})
	<-done
}

func TestTimerRoutineError(t *testing.T) {
	// This value is a joke of "烫"/"쳌".
	// See <https://www.zhihu.com/question/23600507>
	illegalClock := ClockID(0xcccc)
	_, err := illegalClock.NewTimer(time.Second)
	if err == nil {
		t.Fatalf("should not able to create timer on illegal clock")
	}
	t.Logf("err = %v", err)
	// check leak
	if cntChan != LIMIT_TIMER {
		t.Fatalf("channel leaking: cntChan = %d", cntChan)
	}
}
