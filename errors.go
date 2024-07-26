package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func isSame(z, d float64, i int) bool {
	if (z == d || d - z <= 0.000001) && i > 0 {
		return true
	}

	return false
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z, d := 1.0, 0.0
	for i := 0;; i++ {
		z -= (z*z - x) / (2*z)

		if isSame(z,d,i) {
			fmt.Println("We calculated sqrt of x in", i, "times")
			return z, nil
		}

		d = z
	}
}

func main() {
	i, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)	
	} else {
		fmt.Println(i)	
	}
}
