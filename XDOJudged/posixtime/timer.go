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

// +build linux

package posixtime

// We do NOT use timer_* system calls since they rely on OS signals.
// Insteadly we emulate them using ClockID.Sleep and provide an interface
// like time.Timer and time.Ticker.

import (
	"errors"
	"time"
)

// The TimerEvent type represents a timer expiration event.
type TimerEvent struct {
	// If not nil, the timer has encountered a problem and stopped.
	Err error
	// If Err is nil, it contains the clock time of timer expiration.
	// Otherwise is nil.
	Time *time.Time
	// It contains a caller specified value (like sigev_value).
	Value interface{}
}

// The Timer type represents a single event.  When the Timer expires, a
// TimerEvent will be sent on C, unless the Timer was created by
// AfterFunc. A Timer must be created with ClockID.NewTimer or
// ClockID.AfterFunc.
type Timer struct {
	C <-chan TimerEvent

	activeCh chan bool
	clock    ClockID
	handler  func(TimerEvent)
	value    interface{}
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
	active := <-t.activeCh
	t.activeCh <- false
	return active
}

// This internal function set t with empty activeCh, empty C to expire after
// duration d.  If there is an error setting t, the timer expires
// immediately.
func (t *Timer) arm(d time.Duration) {
	t.activeCh <- true
	// Create a goroutine to sleep for duration d and call the handler.
	go func() {
		// set up a TimerEvent and call handler at last.
		var ev TimerEvent

		// Sleep a while.
		err := t.clock.Sleep(d)
		if err != nil {
			ev = TimerEvent{err, nil, t.value}
		}

		// Check if the timer is still active.
		active := <-t.activeCh
		t.activeCh <- false

		// This is active.  We should call the handler.
		if active {
			if err == nil {
				now, err := t.clock.GetTime()
				if err != nil {
					ev = TimerEvent{err, nil, t.value}
				} else {
					ev = TimerEvent{nil, now, t.value}
				}
			}
			t.handler(ev)
		}
	}()
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
func (t *Timer) Reset(d time.Duration) error {
	active := <-t.activeCh

	// Detect an obvious error and report it.
	// But a t.C not drained can not be detected.
	if active {
		t.activeCh <- true
		return errors.New("should not Reset an active timer")
	}

	t.arm(d)
	return nil
}

// NewTimer creates a new Timer on clock.  The Timer would send a
// TimerEvent its channel after at least duration d, or encounter an
// error.
func (clock ClockID) NewTimer(d time.Duration, v interface{}) *Timer {
	ch := make(chan TimerEvent)
	t := Timer{
		C:        ch,
		activeCh: make(chan bool, 1),
		clock:    clock,
		handler: func(ev TimerEvent) {
			ch <- ev
		},
		value: v,
	}

	t.arm(d)
	return &t
}

// AfterFunc waits for the duration to elapse and then calls f in its own
// goroutine, with a TimerEvent as arugment.  It returns a Timer that can
// be used to cancel the call using its Stop method.  If AfterFunc encounter
// an error, f is called immediately in its own goroutine.
func (clock ClockID) AfterFunc(d time.Duration, v interface{},
	f func(TimerEvent)) *Timer {
	ch := make(chan TimerEvent)
	t := Timer{
		C:        ch,
		activeCh: make(chan bool, 1),
		clock:    clock,
		handler:  f,
	}
	t.arm(d)
	return &t
}
