package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, v := range strings.Fields(s) {
		count, ok := result[v]
		if ok {
			result[v] = count + 1
		} else {
			result[v] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
