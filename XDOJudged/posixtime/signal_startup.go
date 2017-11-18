// Block SIGRTMIN on startup.
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

// +build !plan9, !windows

package posixtime

// This function blocks SIGRTMIN before Go runtime startups and creates
// worker threads, to ensure we can handle it with sigwaitinfo.

/*
#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread
#include <errno.h>
#include <signal.h>
#include <stdint.h>
#include <stdlib.h>

int SigmaskInitErrno = -1;

static void init_sigmask(void) __attribute__((constructor));

static void init_sigmask(void)
{
	sigset_t sset;

	if (sigemptyset(&sset) < 0) {
		SigmaskInitErrno = errno;
		return;
	}

	if (sigaddset(&sset, SIGRTMIN) < 0) {
		SigmaskInitErrno = errno;
		return;
	}

	SigmaskInitErrno = pthread_sigmask(SIG_BLOCK, &sset, NULL);
}
*/
import "C"

import (
	"syscall"
)

func init() {
	// If we can't set sigmask, panic at startup.
	if C.SigmaskInitErrno != 0 {
		panic(syscall.Errno(C.SigmaskInitErrno))
	}

	// Create goroutine demultiplexing the incoming SIGRTMINs.
	go demux()
}
