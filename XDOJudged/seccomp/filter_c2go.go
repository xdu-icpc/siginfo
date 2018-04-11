// +build ignore

package seccomp

/*
#include <linux/filter.h>
*/
import "C"

type sockFprog C.struct_sock_fprog
