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

	// It seems ugly but it have to be, since we can't store unsafe
	// pointers.
	_, _, errno := unix.RawSyscall(unix.SYS_SECCOMP,
		1, flags,
		uintptr(unsafe.Pointer(&sockFprog{
			Len:    uint16(len(filter)),
			Filter: uintptr(unsafe.Pointer(&filter[0])),
		})))
	if errno != 0 {
		return errno
	}

	return nil
}
