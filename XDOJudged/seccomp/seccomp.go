// Document of package seccomp.
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

// Package seccomp provides an interface to the Linux Kernel's syscall
// filtering mechanism.  It's API is designed to abstract away the
// underlying BPF based syscall filter language and provide a coventional
// function-call based filtering interface that should be familiar to
// application developers.
//
// Unlike "github.com/seccomp/libseccomp-golang", this package is native
// in Go and does not require libseccomp.
package seccomp
