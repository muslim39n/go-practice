package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var c int

func MutexTest() {
	var mutex sync.Mutex

	go mutexTest(0, &mutex)
	go mutexTest(1, &mutex)
	go mutexTest(2, &mutex)

	// go test(0)
	// go test(1)
	// go test(2)

	time.Sleep(time.Millisecond * 1000)
}

func test(n int) {
	c = 0
	for i := 0; i < 10; i++ {
		c++
		fmt.Printf("func #%d c = %d\n", n, c)
		time.Sleep(time.Millisecond)
	}
}

func mutexTest(n int, mutex *sync.Mutex) {
	mutex.Lock()
	c = 0
	for i := 0; i < 10; i++ {
		c++
		fmt.Printf("func #%d c = %d\n", n, c)
		time.Sleep(time.Millisecond)
	}
	mutex.Unlock()
	for i := 0; i < 10; i++ {
		c++
		fmt.Printf("func #%d i = %d\n", n, i)
		time.Sleep(time.Millisecond)
	}
}
