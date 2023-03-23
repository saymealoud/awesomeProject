package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum int
	//mutex sync.Mutex
	mutex sync.RWMutex
)

func main() {
	run()
}

func run() {
	for i := 0; i < 100; i++ {
		go add(1000)
	}

	for i := 0; i < 10; i++ {
		go fmt.Println("和为:", readSum())
	}

	time.Sleep(8 * time.Second)
}
func readSum() int {
	//只获取读锁
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}
func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
