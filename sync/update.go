package main

import (
	"fmt"
	"sync"
)

func main() {
	run1()
}
func run1() {

	var wg sync.WaitGroup

	//因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			add(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			fmt.Println("和为:", readSum())
		}()
	}

	//一直等待，只要计数器值为0
	wg.Wait()
}
