package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i, _ := range a {
		b := make([]uint8, dx)
		for j, _ := range b {
			b[j] = uint8((i + j) / 2)
		}
		a[i] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}
