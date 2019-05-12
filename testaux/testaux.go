package testaux

/*
#include <signal.h>

const int offset_pid = __builtin_offsetof(siginfo_t, si_pid);
const int offset_uid = __builtin_offsetof(siginfo_t, si_uid);
const int offset_timer = __builtin_offsetof(siginfo_t, si_timerid);
const int offset_overrun = __builtin_offsetof(siginfo_t, si_overrun);
const int offset_status = __builtin_offsetof(siginfo_t, si_status);
const int offset_value = __builtin_offsetof(siginfo_t, si_value);
const int offset_addr = __builtin_offsetof(siginfo_t, si_addr);
const int offset_addr_lsb = __builtin_offsetof(siginfo_t, si_addr_lsb);
const int offset_band = __builtin_offsetof(siginfo_t, si_band);
const int offset_fd = __builtin_offsetof(siginfo_t, si_fd);
const int offset_call_addr = __builtin_offsetof(siginfo_t, si_call_addr);
const int offset_syscall = __builtin_offsetof(siginfo_t, si_syscall);
const int offset_arch = __builtin_offsetof(siginfo_t, si_arch);
const int offset_utime = __builtin_offsetof(siginfo_t, si_utime);
const int offset_stime = __builtin_offsetof(siginfo_t, si_stime);
*/
import "C"

var (
	OffsetPid = uintptr(C.offset_pid)
	OffsetUid = uintptr(C.offset_uid)
	OffsetTimer = uintptr(C.offset_timer)
	OffsetOverrun = uintptr(C.offset_overrun)
	OffsetStatus = uintptr(C.offset_status)
	OffsetValue = uintptr(C.offset_value)
	OffsetAddr = uintptr(C.offset_addr)
	OffsetAddrLsb = uintptr(C.offset_addr_lsb)
	OffsetBand = uintptr(C.offset_band)
	OffsetFd = uintptr(C.offset_fd)
	OffsetCallAddr = uintptr(C.offset_call_addr)
	OffsetSyscall = uintptr(C.offset_syscall)
	OffsetArch = uintptr(C.offset_arch)
	OffsetUtime = uintptr(C.offset_utime)
	OffsetStime = uintptr(C.offset_stime)
)
