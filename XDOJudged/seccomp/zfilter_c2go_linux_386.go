// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs filter_c2go.go

package seccomp

type sockFprog struct {
	Len       uint16
	Pad_cgo_0 [2]byte
	Filter    uintptr
}
