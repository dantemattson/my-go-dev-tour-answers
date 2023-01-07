package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (a [][]uint8) {
	var b []uint8
	for i := 0; i < dy; i++ {
		b = nil
		for j := 0; j < dx; j++ {
			b = append(b, uint8((i*j)*4))
		}
		a = append(a, b)
	}
	return
}

func main() {
	pic.Show(Pic)
}
