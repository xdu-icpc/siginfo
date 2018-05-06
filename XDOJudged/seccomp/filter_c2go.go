// +build ignore

package seccomp

import "golang.org/x/net/bpf"

/*
#include <linux/filter.h>
*/
import "C"

type sockFprog C.struct_sock_fprog
