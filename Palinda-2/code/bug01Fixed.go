package main


import "fmt"


// I want this program to print "Hello world!", but it doesn't work.

func main() {
	ch := make(chan string)
	go func(){
		ch <- "Hello World!"
		close(ch)
	}()
	fmt.Println(<-ch)
}
