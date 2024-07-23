package main

import "fmt"

func fibonacci() func() int {
	n := 0
	a,b := 0,1

	return func() int {
		n++
		if n <= 1 {
			return 0
		}

		a,b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}