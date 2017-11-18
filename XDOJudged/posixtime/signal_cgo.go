// Get macros and constants from signal.h.
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

/*
#include <signal.h>

int _SIGRTMIN;
static void init_sigrtmin(void) __attribute__((constructor));

const int offset_of_si_value = __builtin_offsetof(siginfo_t, si_value);

static void init_sigrtmin(void)
{
	_SIGRTMIN = SIGRTMIN;
}
*/
import "C"

import (
	"syscall"
	"unsafe"
)

const sizeofSigset = C.sizeof_sigset_t
const _SIGEV_NONE = C.SIGEV_NONE
const _SIGEV_SIGNAL = C.SIGEV_SIGNAL
const _SIGEV_THREAD = C.SIGEV_THREAD

var _SIGRTMIN = C._SIGRTMIN

type sigset C.sigset_t
type siginfo C.siginfo_t
type sigevent C.struct_sigevent

// This signal is used by the package.  Do not handle it with os/signal.
// This is implemented as a function instead if a const, because in NPTL
// environments SIGRTMIN is not a real compile-time constant.
func SIGRTMIN() syscall.Signal {
	return syscall.Signal(_SIGRTMIN)
}

func sigsetRTMIN() sigset {
	var ret sigset
	C.sigemptyset((*C.sigset_t)(unsafe.Pointer(&ret)))
	C.sigaddset((*C.sigset_t)(unsafe.Pointer(&ret)), _SIGRTMIN)
	return ret
}

func (i *siginfo) getValue() uintptr {
	// Some platform has si_value in anonymous union, so we have to use
	// a dirty expression.
	return *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(i)) +
		uintptr(C.offset_of_si_value)))
}

func (e *sigevent) setValue(v uintptr) {
	*(*uintptr)(unsafe.Pointer(&e.sigev_value)) = v
}
