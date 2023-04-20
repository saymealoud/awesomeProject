package main

import (
	"fmt"
	"unsafe"
)

/**
1.先使用 new 函数声明一个 *person 类型的指针变量 p。
2.然后把 *person 类型的指针变量 p 通过 unsafe.Pointer，转换为 *string 类型的指针变量 pName。
3.因为 person 这个结构体的第一个字段就是 string 类型的 Name，所以 pName 这个指针就指向 Name 字段（偏移为 0），对 pName 进行修改其实就是修改字段 Name 的值。
4.因为 Age 字段不是 person 的第一个字段，要修改它必须要进行指针偏移运算。所以需要先把指针变量 p 通过 unsafe.Pointer 转换为 uintptr，这样才能进行地址运算。既然要进行指针偏移，那么要偏移多少呢？
	这个偏移量可以通过函数 unsafe.Offsetof 计算出来，该函数返回的是一个 uintptr 类型的偏移量，有了这个偏移量就可以通过 + 号运算符获得正确的 Age 字段的内存地址了，也就是通过 unsafe.Pointer 转换后的 *int 类型的指针变量 pAge。
5.然后需要注意的是，如果要进行指针运算，要先通过 unsafe.Pointer 转换为 uintptr 类型的指针。指针运算完毕后，还要通过 unsafe.Pointer 转换为真实的指针类型（比如示例中的 *int 类型），这样可以对这块内存进行赋值或取值操作。
6.有了指向字段 Age 的指针变量 pAge，就可以对其进行赋值操作，修改字段 Age 的值了。
*/

func main() {

	p := new(unitptr_person)

	//Name是person的第一个字段不用偏移，即可通过指针修改

	pName := (*string)(unsafe.Pointer(p))

	*pName = "飞雪无情"

	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改

	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))

	*pAge = 20

	pPhone := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.phone)))

	*pPhone = "17629299160"

	fmt.Println(*p)

}

type unitptr_person struct {
	Name string

	Age int

	phone string
}
