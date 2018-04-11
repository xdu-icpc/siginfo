package seccomp

/*
#include <sys/prctl.h>

int PrctlSetNoNewPrivsErrno = -1;

static void init_set_no_new_privs(void) __attribute__((constructor));

static void init_set_no_new_privs(void)
{
	PrctlSetNoNewPrivsErrno = prctl(PR_SET_NO_NEW_PRIVS, 1, 0, 0, 0);
}
*/
import "C"

import "syscall"

func init() {
	if C.PrctlSetNoNewPrivsErrno != 0 {
		panic(syscall.Errno(C.PrctlSetNoNewPrivsErrno))
	}
}
