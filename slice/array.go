package main

import "fmt"

func main() {

	ss := []string{"飞雪无情", "张三"}

	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss))

	ss = append(ss, "李四", "王五")

	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss))

	fmt.Println(ss)
}
