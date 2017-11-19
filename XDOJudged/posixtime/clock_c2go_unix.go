// POSIX specified types and constants.
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

// This file must be translated by c2go.sh for platforms.
// +build ignore

package posixtime

/*
#include <time.h>
*/
import "C"

// System clock ID type
type ClockID C.clockid_t

// Clocks defined by IEEE 1003.1-2008
const (
	// System-wide realtime clock
	CLOCK_REALTIME ClockID = C.CLOCK_REALTIME

	// Monotonic system-wide clock
	CLOCK_MONOTONIC ClockID = C.CLOCK_MONOTONIC

	// High-resolution timer from the CPU
	CLOCK_PROCESS_CPUTIME_ID ClockID = C.CLOCK_PROCESS_CPUTIME_ID

	// Thread-specific CPU-time clock
	CLOCK_THREAD_CPUTIME_ID ClockID = C.CLOCK_THREAD_CPUTIME_ID
)

const _TIMER_ABSTIME = C.TIMER_ABSTIME

var posixClocks = [...]ClockID{CLOCK_REALTIME, CLOCK_MONOTONIC,
	CLOCK_PROCESS_CPUTIME_ID, CLOCK_THREAD_CPUTIME_ID}
