package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)

func Select() {
	var (
		cha    = make(chan int)
		chb    = make(chan int)
		chc    = make(chan int)
		chquit = make(chan bool)
	)

	go selectA(cha, chquit)
	go selectB(chb, chquit)
	go selectC(chc, chquit)

	for i := 1; i < 100; i++ {
		fmt.Printf("> %d", i)
		var x int
		select {
		case x = <-cha:
			fmt.Printf("  A: %d\n", x)
		case x = <-chb:
			fmt.Printf("  B: %d\n", x)
		case x = <-chc:
			fmt.Printf("  c: %d\n", x)
		}
	}
	chquit <- true
	chquit <- true
	chquit <- true
	time.Sleep(50 * time.Millisecond)
}

func selectA(ch chan int, quit chan bool) {
	a := rand.Int()
	for {
		select {
		case <-quit:
			fmt.Println("Close A")
			break
		case ch <- a:
			a = rand.Int()
		}
	}
}

func selectB(ch chan int, quit chan bool) {
	a := rand.Int()
	for {
		select {
		case <-quit:
			fmt.Println("Close B")
			break
		case ch <- a:
			a = rand.Int()
		}
	}
}

func selectC(ch chan int, quit chan bool) {
	a := rand.Int()
	for {
		select {
		case <-quit:
			fmt.Println("Close C")
			break
		case ch <- a:
			a = rand.Int()
		}
	}
}
