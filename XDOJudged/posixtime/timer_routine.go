// Low level routine implementing key timer functions.
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
	"unsafe"

	"golang.org/x/sys/unix"
)

var zeroits = itimerspec{}

func timerRoutine(c ClockID, ts unix.Timespec, chstop chan struct{},
	callback func()) error {
	lockUselessLock()
	chexpire := make(chan struct{})
	unlockUselessLock()

	// create the timer
	sigev := sigevent{
		sigev_notify: _SIGEV_SIGNAL,
		sigev_signo:  _SIGRTMIN,
	}
	// XXX we are storing the address of a channel into KERNEL.  It's
	// absolutely _unsafe_.  Make sure GC won't destroy the channel too
	// early.
	sigev.setValue(uintptr(unsafe.Pointer(&chexpire)))

	t, err := timerCreate(c, &sigev)
	if err != nil {
		return err
	}

	// set time for the timer
	its := itimerspec{
		interval: unix.Timespec{},
		value:    ts,
	}
	_, err = timerSetTime(t, 0, &its)
	if err != nil {
		timerDelete(t) // do not leak kernel object
		return err
	}

	go func() {
		select {
		case <-chstop:
			// stop the timer by setting a zero itimerspec.
			its, err := timerSetTime(t, 0, &zeroits)
			if err != nil {
				panic(err)
			}

			if *its == zeroits {
				// This means the timer has already expired.
				// There is a potential race conditon: the expiration
				// signal is still in queue and chexpire is NOT closed
				// yet.  Then we can't give chexpire to GC, or demux()
				// may close a destroyed channel.
				select {
				case <-chexpire: // Now we know chexpire is closed.
				}
			} else {
				// Otherwise, the timer has been stopped successfully
				// before it could expire.  Though we can let GC close it,
				// we do this ourselves to make this explicitly.
				close(chexpire)
			}
		case <-chexpire:
			// The timer has expired and chexpire has been closed.
		}

		// Callback is called unconditionally.  The caller should implement
		// cancel support.
		go callback()
		err = timerDelete(t)
	}()

	return nil
}
