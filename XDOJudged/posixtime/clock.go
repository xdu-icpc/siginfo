// Manipulate POSIX process CPU-time clocks.
// Copyright (C) 2017  Laboratory of ACM/ICPC, Xidian University

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warramty of
// MERCHANTABILITY or FITNESS FOR A PARICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Author: Xi Ruoyao <ryxi@stu.xidian.edu.cn>

// +build linux

package posixtime

import (
	"syscall"
	"time"
	"unsafe"
)

// In Linux clockid_t is int.
type ClockID int

// Pre-defined clocks in Linux.
const (
	// System-wide realtime clock
	CLOCK_REALTIME ClockID = iota

	// Monotonic system-wide clock
	CLOCK_MONOTONIC

	// High-resolution timer from the CPU
	CLOCK_PROCESS_CPUTIME_ID

	// Thread-specific CPU-time clock
	CLOCK_THREAD_CPUTIME_ID

	// Monotonic system-wide clock, not adjusted for frequency scaling.
	CLOCK_MONOTONIC_RAW

	// System-wide realtime clock, updated only on ticks.
	CLOCK_REALTIME_COARSE

	// Monotonic system-wide clock, updated only on ticks.
	CLOCK_MONOTONIC_COARSE

	//Monotonic system wide clock that includes time spent in suspension.
	CLOCK_BOOTTIME

	// Like CLOCK_REALTIME but also wakes suspended system.
	CLOCK_REALTIME_ALARM

	// Like CLOCK_BOOTTIME but also wakes suspended system.
	CLOCK_BOOTTIME_ALARM

	_ // No such clock.

	// System-wide realtime clock using International Atomic Time.
	CLOCK_TAI

	CLOCK_PREDEF_NUM int = iota // The number of predefined clocks.
)

// GetRes returns resolution (precision) of a POSIX clock.
func (clock ClockID) GetRes() (*time.Time, error) {
	var ts syscall.Timespec

	_, _, errno := syscall.Syscall(syscall.SYS_CLOCK_GETRES,
		uintptr(clock),
		uintptr(unsafe.Pointer(&ts)),
		uintptr(0))
	if errno != 0 {
		return nil, errno
	}

	ret := time.Unix(ts.Unix())
	return &ret, nil
}

// Returns a ClockID of a POSIX CPU-time clock of the given PID.
//
// Note: a CPU-time clock is bound to a PID, not a specific process.
// If a new process assumed the PID, the clock would show the CPU
// time of this new process.
func GetCPUClockID(pid int) (*ClockID, error) {
	/* This magic expression is from Linux kernel ABI for CPU clocks.  */
	id := ClockID((^pid)<<3 | 2)

	/* Do a clock_getres call to validate it.  */
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

// Sleep pauses the current goroutine for at least duration d on a POSIX
// clock.  A negative or zero duration causes Sleep to return immediately.
func (clock ClockID) Sleep(d time.Duration) error {
	// TODO: POSIX said this should be a cancellation point.  But I don't
	// know how to do it.

	// The system call doesn't support predefined CPU clocks.
	// Special case them.

	if clock == CLOCK_THREAD_CPUTIME_ID {
		return syscall.EINVAL
	}
	if clock == CLOCK_PROCESS_CPUTIME_ID {
		clock = ClockID((^0)<<3 | 2)
	}

	ts := syscall.NsecToTimespec(d.Nanoseconds())

	_, _, errno := syscall.Syscall6(syscall.SYS_CLOCK_NANOSLEEP,
		uintptr(clock), uintptr(0), uintptr(unsafe.Pointer(&ts)),
		uintptr(0), uintptr(0), uintptr(0))

	if errno == 0 {
		return nil
	}

	return errno
}
