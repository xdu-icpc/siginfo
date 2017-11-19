// Linux specified constants.
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

// Pre-defined clocks in Linux.
const (
	// Monotonic system-wide clock, not adjusted for frequency scaling.
	CLOCK_MONOTONIC_RAW ClockID = C.CLOCK_MONOTONIC_RAW

	// System-wide realtime clock, updated only on ticks.
	CLOCK_REALTIME_COARSE ClockID = C.CLOCK_REALTIME_COARSE

	// Monotonic system-wide clock, updated only on ticks.
	CLOCK_MONOTONIC_COARSE ClockID = C.CLOCK_MONOTONIC_COARSE

	//Monotonic system wide clock that includes time spent in suspension.
	CLOCK_BOOTTIME ClockID = C.CLOCK_BOOTTIME

	// Like CLOCK_REALTIME but also wakes suspended system.
	CLOCK_REALTIME_ALARM ClockID = C.CLOCK_REALTIME_ALARM

	// Like CLOCK_BOOTTIME but also wakes suspended system.
	CLOCK_BOOTTIME_ALARM ClockID = C.CLOCK_BOOTTIME_ALARM

	// System-wide realtime clock using International Atomic Time.
	CLOCK_TAI ClockID = C.CLOCK_TAI
)

var platformClocks = [...]ClockID{CLOCK_MONOTONIC_RAW,
	CLOCK_REALTIME_COARSE, CLOCK_BOOTTIME, CLOCK_REALTIME_ALARM,
	CLOCK_BOOTTIME_ALARM, CLOCK_TAI}
