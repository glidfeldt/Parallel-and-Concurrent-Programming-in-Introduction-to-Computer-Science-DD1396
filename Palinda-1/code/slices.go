package main


import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	slices := make([][]uint8, dy)
	for y := range slices {
		slices[y] = make([]uint8, dx)
	}
	for y, slice := range slices{
		for x := range slice {
			slice[x] = uint8(x^y)
		}
	}
	return slices
}

func main() {
	pic.Show(Pic)
}
