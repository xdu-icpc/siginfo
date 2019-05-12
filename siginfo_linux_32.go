//SPDX-License-Identifier: Beerware

// +build arm 386

package siginfo

type SiginfoHeader struct {
	Signo, Errno, Code int32
}
