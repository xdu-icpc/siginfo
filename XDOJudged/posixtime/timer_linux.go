// Wrap POSIX timer syscalls for Linux.
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
	"runtime"
	"unsafe"

	"golang.org/x/sys/unix"
)

type timerid uintptr

func timerCreate(c ClockID, e *sigevent) (timerid, error) {
	if e.sigev_notify == _SIGEV_THREAD {
		// this is not supported now
		return 0, unix.EINVAL
	}

	ret := timerid(0)

	for {
		_, _, errno := unix.RawSyscall(unix.SYS_TIMER_CREATE, uintptr(c),
			uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(&ret)))

		if errno == 0 {
			return ret, nil
		}

		if errno == unix.EAGAIN {
			continue
		}

		if errno == unix.ENOMEM {
			// if lucky we can find some memory
			runtime.GC()
			continue
		}

		return ret, errno
	}
}

type itimerspec struct {
	interval, value unix.Timespec
}

func timerSetTime(t timerid, f int, s *itimerspec) (*itimerspec, error) {
	oldspec := itimerspec{}

	_, _, errno := unix.RawSyscall6(unix.SYS_TIMER_SETTIME, uintptr(t),
		uintptr(f), uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(&oldspec)), uintptr(0), uintptr(0))
	if errno != 0 {
		return nil, errno
	}

	return &oldspec, nil
}

func timerDelete(t timerid) error {
	_, _, errno := unix.RawSyscall(unix.SYS_TIMER_DELETE, uintptr(t),
		uintptr(0), uintptr(0))
	if errno != 0 {
		return errno
	}

	return nil
}
