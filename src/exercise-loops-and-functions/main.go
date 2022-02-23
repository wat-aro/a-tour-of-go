package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		y := z
		z -= (z*z - x) / (2 * z)

		if math.Abs(1-y/z) < 0.001 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Print(math.Sqrt(2))
}
