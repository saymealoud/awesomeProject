package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 3

	iv := reflect.ValueOf(i)

	it := reflect.TypeOf(i)

	fmt.Println(iv, it) //3 int

	//reflect.Value to int

	i1 := iv.Interface().(int)

	fmt.Println(i1)

	// 传入变量的指针才可以。 因为传递的是一个指针，所以需要调用 Elem 方法找到这个指针指向的值，这样才能修改。 最后我们就可以使用 SetInt 方法修改值了。
	ipv := reflect.ValueOf(&i)

	ipv.Elem().SetInt(4)

	fmt.Println(i)

}
