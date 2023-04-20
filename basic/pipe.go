package main

import (
	"fmt"
)

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	var t1 int
	t1 = <-pipe
	pipe <- 4

	fmt.Println(len(pipe))
	fmt.Println(t1)

}
func main() {
	test_pipe()
}
