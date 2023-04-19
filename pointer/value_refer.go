package main

import "fmt"

/**
Go 语言中的函数传参都是值传递
除了 struct 外，还有浮点型、整型、字符串、布尔、数组，这些都是值类型。

&xx  值传递的是指针，也是内存地址。通过内存地址可以找到原数据的那块内存，所以修改它也就等于修改了原数据。

*/

type address struct {
	province string
	city     string
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func main() {
	//add := address{province: "北京", city: "北京"}
	//printString(add)
	//printString(&add)

	var si fmt.Stringer = address{province: "上海", city: "上海"}
	printString(si)
	sip := &si
	printString(*sip)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}
