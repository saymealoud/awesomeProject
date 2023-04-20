package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**

切片的高效还体现在 for range 循环中，因为循环得到的临时变量也是个值拷贝，所以在遍历大的数组时，切片的效率更高。
*/

func main() {

	a1 := [2]string{"飞雪无情", "张三"}

	fmt.Printf("函数main数组指针：%p\n", &a1)

	arrayF(a1)

	s1 := a1[0:1]

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)

	sliceF(s1)

}

func arrayF(a [2]string) {

	fmt.Printf("函数arrayF数组指针：%p\n", &a)

}

func sliceF(s []string) {

	fmt.Printf("函数sliceF Data：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)

}
