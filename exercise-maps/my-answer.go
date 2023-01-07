package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) (a map[string]int) {
	a = make(map[string]int)
	for _, v := range strings.Fields(s) {
		val, exists := a[v]
		if exists {
			a[v] = val + 1
		} else {
			a[v] = 1
		}
	}
	return
}

func main() {
	wc.Test(WordCount)
}
