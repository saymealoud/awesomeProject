package main

import (
	"fmt"
	"time"
)

func test_goroute(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func test_print(a int) {
	fmt.Println(a)
}
func main() {
	//test_goroute(100, 200)
	for i := 0; i < 100; i++ {
		test_print(i)
	}

	time.Sleep(time.Second)
}
