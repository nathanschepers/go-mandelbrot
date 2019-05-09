package main

import (
	"math/cmplx"
)

func inMandel(c complex128, iterations int) (bool, int) {
	z := 0 + 0i
	lastZ := z

	for i := 0; i < iterations; i++ {
		z = cmplx.Pow(lastZ, 2) + c
		if cmplx.Abs(z) > 2 {
			return false, i
		}
		lastZ = z
	}
	return true, iterations
}
