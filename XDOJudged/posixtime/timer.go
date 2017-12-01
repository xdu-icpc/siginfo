// Manipulate timers on POSIX clocks.
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

package posixtime

import (
	"errors"
	"sync/atomic"
	"time"

	"golang.org/x/sys/unix"
)

// The TimerEvent type represents a timer expiration event.
type TimerEvent struct {
	// It contains the clock time of timer expiration.
	Time time.Time
}

// The Timer type represents a single event.  When the Timer expires, a
// TimerEvent will be sent on C, unless the Timer was created by
// AfterFunc. A Timer must be created with ClockID.NewTimer or
// ClockID.AfterFunc.
type Timer struct {
	C <-chan TimerEvent

	clock  ClockID
	armed  int32
	stopch chan<- struct{}
	handle func()
}

// Stop prevents the Timer from firing.  It returns true if the call stops
// the timer, false if the timer has already expired or been stopped.  Stop
// does not close the channel, to prevent a read from the channel succeeding
// incorrectly.
//
// To prevent a timer created with ClockID.NewTimer from firing after a call
// to Stop, check the return value and drain the channel.  Foe example,
// assuming the program has not received from t.C already:
//
//     if !t.Stop() {
//         <-t.C
//     }
//
// This cannot be done concurrent to other receives from the Timer's
// channel.
//
// For a timer created with ClockID.AfterFunc(d, f), it t.Stop returns
// false, then the timer has already expired and the function f has been
// started in its own goroutine; Stop does not wait for f to complete
// before returning.  If the caller needs to know whether f is completed,
// it must coordinate with f explicitly.
func (t *Timer) Stop() bool {
	active := atomic.SwapInt32(&t.armed, 0)
	if active == 1 {
		close(t.stopch)
		return true
	}
	return false
}

// This internal function set _unarmed_ t with empty C to expire after
// duration d.  If there is an error setting t, this function returns and
// do nothing.
func (t *Timer) arm(d time.Duration) error {
	stopch := make(chan struct{})
	t.stopch = stopch
	t.armed = 1
	err := timerRoutine(t.clock,
		unix.NsecToTimespec(d.Nanoseconds()),
		stopch,
		func() {
			active := atomic.SwapInt32(&t.armed, 0)
			if active == 1 {
				t.handle()
			}
		})
	if err != nil {
		t.armed = 0
	}
	return err
}

// Reset changes the timer to expire after duration d.  It must be used on
// an inactive timer with a drained t.C.  If a program has already received
// a value from t.C, the timer is know to have expired, then t.Reset can be
// used directly.  If not, the timer must be stopped and-if Stop reports
// that the timer expired before being stopped-the channel explicitly
// drained:
//
//     if !t.Stop() {
//         <-t.C
//     }
//     t.Reset(d)
//
// This should not be done concurrent to other receives from the timer's
// channel.
//
// If t is active, Reset would return an error indicating this issue and
// do nothing.  However if t is inactive but t.C is not drained, Reset can
// not know this and there may be severe problems. Do NOT do that.
//
// If Reset can not arm the timer, it return an error and do nothing.
func (t *Timer) Reset(d time.Duration) error {
	// Detect an obvious error and report it.
	// But a t.C not drained can not be detected.
	if atomic.LoadInt32(&t.armed) == 1 {
		return errors.New("should not Reset an active timer")
	}

	return t.arm(d)
}

// NewTimer creates a new Timer on clock.  The Timer would send a
// TimerEvent its channel after at least duration d.
func (clock ClockID) NewTimer(d time.Duration) (*Timer, error) {
	ch := make(chan TimerEvent)
	t := Timer{
		C:     ch,
		armed: 0,
		clock: clock,
		handle: func() {
			t, _ := clock.GetTime()
			ch <- TimerEvent{*t}
		},
	}

	err := t.arm(d)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// AfterFunc waits for the duration to elapse and then calls f in its own
// goroutine, with a TimerEvent as arugment.  It returns a Timer that can
// be used to cancel the call using its Stop method.
func (clock ClockID) AfterFunc(d time.Duration, f func(TimerEvent)) (
	*Timer, error) {
	ch := make(chan TimerEvent)
	t := Timer{
		C:     ch,
		armed: 0,
		clock: clock,
		handle: func() {
			t, _ := clock.GetTime()
			f(TimerEvent{*t})
		},
	}
	err := t.arm(d)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
