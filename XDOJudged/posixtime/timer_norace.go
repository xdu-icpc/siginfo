// Useless functions in production code (see below).
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

// +build !race

package posixtime

// In order to demux the signals from POSIX timers, we store the address
// of a channel into the kernel (struct sigevent), then retrieve it from
// siginfo_t.  Go runtime doesn't know the signal must come after the
// creation of POSIX timer (it's some mysterious system call).  So the
// race detector would believe we are racing.  But actually there is no
// racing.  So in the production code we just do nothing.

// I hope the compiler would eliminate the empty function call.

func lockUselessLock() {
	// do nothing
}

func unlockUselessLock() {
	// do nothing
}
