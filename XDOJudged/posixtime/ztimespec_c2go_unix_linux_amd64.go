// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs timespec_c2go_unix.go

package posixtime

import "golang.org/x/sys/unix"

func timespec(s int64, ns int) unix.Timespec {
	return unix.Timespec{
		Sec:  int64(s),
		Nsec: int64(ns),
	}
}
