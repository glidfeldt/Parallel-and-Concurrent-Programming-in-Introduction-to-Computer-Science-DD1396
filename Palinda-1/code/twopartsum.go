package main

import (
	"fmt"
)
const NUM = 2
// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan <- int) {
	sum := 0
	for _, value := range a {
		sum += value
	}
	res <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9 ,10}
	n := len(a)
	res := make(chan int)
	go Add(a[:n/2], res)
	go Add(a[n/2:], res)
	x, y := <-res, <-res // receive from c
	fmt.Println(x,"+",y,"=",x+y)
}
