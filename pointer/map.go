package main

import "fmt"

/**
func makemap(t *maptype, hint int, h *hmap) *hmap {}

map 修改直接使用的是变量 m 和 p，并没有用到取地址符 &，这是因为它们本来就是指针，所以就没有必要再使用 & 取地址了。

所以在这里，Go 语言通过 make 函数或字面量的包装为我们省去了指针的操作，让我们可以更容易地使用 map。其实就是语法糖，这是编程界的老传统了。

这里的 map 可以理解为引用类型，但是它本质上是个指针，只是可以叫作引用类型而已。在参数传递时，它还是值传递，并不是其他编程语言中所谓的引用传递



引用类型和基本类型
	基本类型一般有初始化默认值
	引用类型初始化 无默认值 一般为空  Java 为 null   Go 则为 nil
*/

func main() {

	m := make(map[string]int)

	m["飞雪无情"] = 18

	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])

	modifyMap(m)

	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])

}

func modifyMap(p map[string]int) {

	p["飞雪无情"] = 20

}
