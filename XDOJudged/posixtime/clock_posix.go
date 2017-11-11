// POSIX specified clock IDs.
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

// +build darwin dragonfly freebsd linux nacl netbsd openbst solaris windows

package posixtime

// Clocks defined by IEEE 1003.1-2008
const (
	// System-wide realtime clock
	CLOCK_REALTIME ClockID = _CLOCK_REALTIME

	// Monotonic system-wide clock
	CLOCK_MONOTONIC ClockID = _CLOCK_MONOTONIC

	// High-resolution timer from the CPU
	CLOCK_PROCESS_CPUTIME_ID ClockID = _CLOCK_PROCESS_CPUTIME_ID

	// Thread-specific CPU-time clock
	CLOCK_THREAD_CPUTIME_ID ClockID = _CLOCK_THREAD_CPUTIME_ID
)
