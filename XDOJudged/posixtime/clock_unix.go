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
// +build dragonfly freebsd linux netbsd openbsd

package posixtime

import (
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

// GetRes returns resolution (precision) of a POSIX clock.
func (clock ClockID) GetRes() (*time.Duration, error) {
	var ts unix.Timespec

	_, _, errno := unix.RawSyscall(unix.SYS_CLOCK_GETRES,
		uintptr(clock),
		uintptr(unsafe.Pointer(&ts)),
		uintptr(0))
	if errno != 0 {
		return nil, errno
	}

	ret := time.Duration(ts.Nano()) * time.Nanosecond
	return &ret, nil
}

// GetTime returns the time of a POSIX clock.
func (clock ClockID) GetTime() (*time.Time, error) {
	var ts unix.Timespec

	_, _, errno := unix.RawSyscall(unix.SYS_CLOCK_GETTIME,
		uintptr(clock),
		uintptr(unsafe.Pointer(&ts)),
		uintptr(0))
	if errno != 0 {
		return nil, errno
	}

	ret := time.Unix(ts.Unix())
	return &ret, nil
}
