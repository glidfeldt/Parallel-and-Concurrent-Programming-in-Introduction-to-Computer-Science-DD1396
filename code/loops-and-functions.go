package main

import (
	"fmt"
	//"math"
)

func Sqrt(x float64) float64 {
	z := 2.5
	i := 0
	var zold float64
	for {
		i+=1
		zold = z
		z -=(z*z-x)/(2*z)
		if z == zold {
			break
		}
	}
	//fmt.Println(i)
	return z
}

func main() {
	fmt.Println(Sqrt(6))
	//fmt.Println(math.Sqrt(6))
}

