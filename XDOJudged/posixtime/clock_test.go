// Unitest of clock_*.go.
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

package posixtime_test

import (
	"testing"
	"time"

	"linux.xidian.edu.cn/git/XDU_ACM_ICPC/XDOJ-next/XDOJudged/posixtime"
)

func TestGetTime(t *testing.T) {
	for _, i := range posixtime.ALL_CLOCKS {
		if i == -1 {
			continue // no such clock
		}
		clock := posixtime.ClockID(i)
		time, err := clock.GetTime()
		if err != nil {
			t.Errorf("can not get the time of clock ID %d: %v", i, err)
		} else {
			t.Logf("clock ID = %d, result = %v", i, time)
		}
	}
}

func TestGetRes(t *testing.T) {
	for _, i := range posixtime.ALL_CLOCKS {
		if i == -1 {
			continue // no such clock
		}
		clock := posixtime.ClockID(i)
		time, err := clock.GetRes()
		if err != nil {
			t.Errorf("can not get the time of clock ID %d: %v", i, err)
		} else {
			t.Logf("clock ID = %d, result = %v", i, time)
		}
	}
}

func TestCPUClock(t *testing.T) {
	clock, err := posixtime.GetCPUClockID(1)
	if err != nil {
		t.Fatalf("can not create CPU-time clock of process 1: %v", err)
	}

	time, err := clock.GetTime()
	if err != nil {
		t.Fatalf("can not get the time of the CPU-time clock: %v", err)
	}

	t.Logf("result = %v", time)
}

func sleepWell(d time.Duration) error {
	t0 := time.Now()
	err := posixtime.CLOCK_MONOTONIC.Sleep(d)
	if err != nil {
		return err
	}

	d1 := time.Now().Sub(t0)
	return checkDuration(d1, d)
}

func TestSleep(t *testing.T) {
	err := sleepWell(sleepDuration)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestWaitUntil(t *testing.T) {
	t0, err := posixtime.CLOCK_MONOTONIC.GetTime()
	if err != nil {
		t.Fatalf("can not get time of CLOCK_MONOTONIC: %v", err)
	}

	ts := t0.Add(sleepDuration)

	err = posixtime.CLOCK_MONOTONIC.WaitUntil(ts)
	if err != nil {
		t.Fatalf("can not wait until a time: %v", err)
	}

	t1, err := posixtime.CLOCK_MONOTONIC.GetTime()
	if err != nil {
		t.Fatalf("can not get time of CLOCK_MONOTONIC: %v", err)
	}

	d := t1.Sub(*t0)
	err = checkDuration(d, sleepDuration)
	if err != nil {
		t.Fatal(err)
	}
}
