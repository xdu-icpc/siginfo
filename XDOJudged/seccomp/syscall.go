package seccomp

import (
	"golang.org/x/net/bpf"
	"golang.org/x/sys/unix"
	"unsafe"
)

const (
	SECCOMP_FILTER_FLAG_TSYNC = 1
	SECCOMP_FILTER_FLAG_LOG   = 2
)

// I don't think "strict mode" can support Go program.

// Install a seccomp filter.
func SeccompFilter(flags uintptr, filter []bpf.RawInstruction) error {
	if len(filter) > 255 {
		return unix.EINVAL
	}

	sockProg := sockFprog{
		Len:    uint16(len(filter)),
		Filter: &filter[0],
	}

	_, _, errno := unix.RawSyscall(unix.SYS_SECCOMP, 1, flags,
		uintptr(unsafe.Pointer(&sockProg)))
	if errno != 0 {
		return errno
	}

	return nil
}
