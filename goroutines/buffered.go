package goroutines

import "fmt"

func Buffered() {
	// c := make(chan int)

	// c <- 123
	// c <- 23
	// fmt.Println(<-c, <-c)
	c := make(chan int, 2)

	c <- 123
	c <- 23
	fmt.Println(<-c, <-c)
}
