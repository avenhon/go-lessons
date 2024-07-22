package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	str := strings.Fields(s)
	for _, key := range str {
		v, ok := m[key]
		if ok {
			v++
			m[key] = v
		} else {
			m[key] = 1	
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}