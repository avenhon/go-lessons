package main

import (
	"fmt"
	"math"
)

func isSame(z, d float64, i int) bool {
	fmt.Println(z, d, i)

	if (z == d || d - z <= 0.000001) && i > 0 {
		return true
	}

	return false
}

func Sqrt(x float64) float64 {
	z, d := 1.0, 0.0
	for i := 0;; i++ {
		z -= (z*z - x) / (2*z)

		if isSame(z,d,i) {
			fmt.Println("We calculated sqrt of x in", i, "times")
			return z
		}

		d = z
	}
}

func main() {
	// Check if our function close or the same with math.Sqrt standart library
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
