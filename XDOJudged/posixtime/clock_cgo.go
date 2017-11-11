// Get macros in time.h.
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

/*
#include <time.h>

#ifndef CLOCK_MONOTONIC
#define CLOCK_MONOTONIC (-1)
#endif

#ifndef CLOCK_PROCESS_CPUTIME_ID
#define CLOCK_PROCESS_CPUTIME_ID (-1)
#endif

#ifndef CLOCK_REALTIME
#define CLOCK_REALTIME (-1)
#endif

#ifndef CLOCK_THREAD_CPUTIME_ID
#define CLOCK_THREAD_CPUTIME_ID (-1)
#endif

#ifndef CLOCK_MONOTONIC_RAW
#define CLOCK_MONOTONIC_RAW (-1)
#endif

#ifndef CLOCK_REALTIME_COARSE
#define CLOCK_REALTIME_COARSE (-1)
#endif

#ifndef CLOCK_MONOTONIC_COARSE
#define CLOCK_MONOTONIC_COARSE (-1)
#endif

#ifndef CLOCK_BOOTTIME
#define CLOCK_BOOTTIME (-1)
#endif

#ifndef CLOCK_REALTIME_ALARM
#define CLOCK_REALTIME_ALARM (-1)
#endif

#ifndef CLOCK_BOOTTIME_ALARM
#define CLOCK_BOOTTIME_ALARM (-1)
#endif

#ifndef CLOCK_TAI
#define CLOCK_TAI (-1)
#endif

#ifndef TIMER_ABSTIME
#define TIMER_ABSTIME (-1)
#endif
*/
import "C"

const (
	_CLOCK_MONOTONIC          = C.CLOCK_MONOTONIC
	_CLOCK_PROCESS_CPUTIME_ID = C.CLOCK_PROCESS_CPUTIME_ID
	_CLOCK_REALTIME           = C.CLOCK_REALTIME
	_CLOCK_THREAD_CPUTIME_ID  = C.CLOCK_THREAD_CPUTIME_ID
	_CLOCK_MONOTONIC_RAW      = C.CLOCK_MONOTONIC_RAW
	_CLOCK_REALTIME_COARSE    = C.CLOCK_REALTIME_COARSE
	_CLOCK_MONOTONIC_COARSE   = C.CLOCK_MONOTONIC_COARSE
	_CLOCK_BOOTTIME           = C.CLOCK_BOOTTIME
	_CLOCK_REALTIME_ALARM     = C.CLOCK_REALTIME_ALARM
	_CLOCK_BOOTTIME_ALARM     = C.CLOCK_BOOTTIME_ALARM
	_CLOCK_TAI                = C.CLOCK_TAI
	_TIMER_ABSTIME            = C.TIMER_ABSTIME
)

// this would be exported for test suite.
var _ALL_CLOCKS = [...]int{_CLOCK_MONOTONIC, _CLOCK_PROCESS_CPUTIME_ID,
	_CLOCK_REALTIME, _CLOCK_THREAD_CPUTIME_ID, _CLOCK_MONOTONIC_RAW,
	_CLOCK_REALTIME_COARSE, _CLOCK_MONOTONIC_COARSE, _CLOCK_BOOTTIME,
	_CLOCK_REALTIME_ALARM, _CLOCK_BOOTTIME_ALARM, _CLOCK_TAI}
