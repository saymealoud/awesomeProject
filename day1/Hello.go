package main

import (
	"fmt"
)

func add(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func main() {
	fmt.Println("hello world")
	fmt.Print(add(2, 4))
}
