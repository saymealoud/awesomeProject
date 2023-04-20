package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
 地址不同了

runtime.stringtoslicebyte 和 runtime.slicebytetostring 这两个函数的源代码
*/

func main() {

	s := "飞雪无情"

	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)

	b := []byte(s)

	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)

	s3 := string(b)

	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
}
