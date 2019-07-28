package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, dy)
	for i := 0; i < dy; i++{
		x[i] = make([]uint8, dx)
	}
	return x
}

func main() {
	pic.Show(Pic)
}
