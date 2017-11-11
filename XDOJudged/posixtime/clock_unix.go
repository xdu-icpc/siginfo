// Manipulate POSIX clocks.
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

// TODO: how about windows?
// +build darwin dragonfly freebsd linux nacl netbsd openbst solaris

package posixtime

import (
	"syscall"
	"time"
	"unsafe"
)

// In Linux clockid_t is int.
type ClockID int

// GetRes returns resolution (precision) of a POSIX clock.
func (clock ClockID) GetRes() (*time.Duration, error) {
	var ts syscall.Timespec

	_, _, errno := syscall.Syscall(syscall.SYS_CLOCK_GETRES,
		uintptr(clock),
		uintptr(unsafe.Pointer(&ts)),
		uintptr(0))
	if errno != 0 {
		return nil, errno
	}

	ret := time.Duration(ts.Nano()) * time.Nanosecond
	return &ret, nil
}

// Returns a ClockID of a POSIX CPU-time clock of the given PID.
//
// Note: a CPU-time clock is bound to a PID, not a specific process.
// If a new process assumed the PID, the clock would show the CPU
// time of this new process.
func GetCPUClockID(pid int) (*ClockID, error) {
	// This magic expression is from Linux kernel ABI for CPU clocks.
	id := ClockID((^pid)<<3 | 2)

	// Do a clock_getres call to validate it.
	_, err := id.GetRes()
	if err == syscall.EINVAL {
		err = syscall.ESRCH
	}

	if err != nil {
		return nil, err
	}

	return &id, nil
}

// GetTime returns the time of a POSIX clock.
func (clock ClockID) GetTime() (*time.Time, error) {
	var ts syscall.Timespec

	_, _, errno := syscall.Syscall(syscall.SYS_CLOCK_GETTIME,
		uintptr(clock),
		uintptr(unsafe.Pointer(&ts)),
		uintptr(0))
	if errno != 0 {
		return nil, errno
	}

	ret := time.Unix(ts.Unix())
	return &ret, nil
}

func (clock ClockID) nanosleep(ts syscall.Timespec, flag int) error {
	// POSIX said clock_nanosleep should be a cancellation point.  But
	// goroutines can't be canceled so we ignore the cancellation things.
	// And also, POSIX said clock_nanosleep can be interrupted and return
	// the remain unslept duration.  Normal goroutines won't handle os
	// signals so we don't use this feature.

	// The system call doesn't support predefined CPU clocks.
	// Special case them.
	if clock == CLOCK_THREAD_CPUTIME_ID {
		return syscall.EINVAL
	}
	if clock == CLOCK_PROCESS_CPUTIME_ID {
		clock = ClockID((^0)<<3 | 2)
	}

	// Do real system call.
	_, _, errno := syscall.Syscall6(syscall.SYS_CLOCK_NANOSLEEP,
		uintptr(clock), uintptr(flag), uintptr(unsafe.Pointer(&ts)),
		uintptr(0), uintptr(0), uintptr(0))

	if errno == 0 {
		return nil
	}

	return errno
}

// Sleep pauses the current goroutine for at least duration d on a POSIX
// clock.  A negative or zero duration causes Sleep to return immediately.
func (clock ClockID) Sleep(d time.Duration) error {
	return clock.nanosleep(syscall.NsecToTimespec(d.Nanoseconds()), 0)
}

// We won't expose this out of the package since Go has different types
// for time and duration.  We'll use different method names and parameter
// types to distinguish durations and absolute times.

// WaitUntil pauses the current goroutine until the POSIX clock reaches
// time t.  A time before current time causes WaitUntil to return
// immediately.
func (clock ClockID) WaitUntil(t time.Time) error {
	ts := syscall.Timespec{
		syscall.Timespec_sec_t(t.Unix()),
		syscall.Timespec_nsec_t(t.Nanosecond()),
	}
	return clock.nanosleep(ts, _TIMER_ABSTIME)
}
