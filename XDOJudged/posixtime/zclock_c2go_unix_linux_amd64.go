// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs clock_c2go_unix.go

package posixtime

type ClockID int32

const (
	CLOCK_REALTIME ClockID = 0x0

	CLOCK_MONOTONIC ClockID = 0x1

	CLOCK_PROCESS_CPUTIME_ID ClockID = 0x2

	CLOCK_THREAD_CPUTIME_ID ClockID = 0x3
)

const _TIMER_ABSTIME = 0x1

var posixClocks = [...]ClockID{CLOCK_REALTIME, CLOCK_MONOTONIC,
	CLOCK_PROCESS_CPUTIME_ID, CLOCK_THREAD_CPUTIME_ID}
