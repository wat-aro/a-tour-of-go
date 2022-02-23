package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		z := x
		x, y = y, x+y
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println((f()))
	}
}
