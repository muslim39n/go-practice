package goroutines

import (
	"fmt"
	"math"
	"math/rand"
)

func FindSquares() {
	var (
		arr [1000]int
		// c   = make(chan int, 1000)
		c = make(chan int)
	)

	for i := range arr {
		arr[i] = rand.Intn(1000)
	}

	go findSquares(arr, c)

	// for {
	// 	v, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	for v := range c {
		fmt.Println(v)
	}
}

func findSquares(arr [1000]int, c chan int) {
	for _, v := range arr {
		root := int(math.Sqrt(float64(v)))
		if root*root == v {
			fmt.Println("put ", v)
			c <- v
		}
	}
	close(c)
}
