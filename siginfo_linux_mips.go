//SPDX-License-Identifier: Beerware

package siginfo

type SiginfoHeader struct {
	Signo, Code, Errno int32
}
