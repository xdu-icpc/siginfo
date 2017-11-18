// Demultiplex POSIX timer's signals.
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
	"syscall"
	"unsafe"
)

func demux() {
	var info siginfo

	mask := sigsetRTMIN()
	for {
		err := sigwaitinfo(&mask, &info)
		if err == syscall.EINTR {
			// If unlucky a signal other than SIGRTMIN may be delivered
			// to our thread.  In this case just do next loop.
			continue
		}
		if err != nil {
			panic(err)
		}

		// XXX this is really _unsafe_.  The receiver should prevent the
		// channel to be destructed by GC.
		ch := *(*chan struct{})(unsafe.Pointer(info.getValue()))
		close(ch)
	}
}
