// Helper to create unix.Timespec.
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

import "golang.org/x/sys/unix"

/*
#include <time.h>
*/
import "C"

// after c2go.sh translation, this function would contain proper type
// conversion to create a Timespec.
func timespec(s int64, ns int) unix.Timespec {
	return unix.Timespec{
		Sec:  C.time_t(s),
		Nsec: C.long(ns),
	}
}
