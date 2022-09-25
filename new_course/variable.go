package main

import "fmt"

var g = "全局变量"

//gl :="全局变量 -端变量声明法   错误写法"

func getUserInfo() (string, int) {
	return "zhangsan", 10
}
func main() {

	//fmt.Println(g)
	//fmt.Println(gl)

	//fmt.Println("声明变量")

	//var username string
	//fmt.Println(username) //声明变量后没有初始化的话    值为空
	//
	//var a1 = "张三"
	//fmt.Println(a1)
	//
	//var m_ = "李四"
	//fmt.Println(m_)

	/*
		var a1, a2 string
		a1 = "aaa"
		a2 = "bbb"
		fmt.Println(a1, a2)
	*/

	/*
		var (

			username string
			age      int
			sex      string
		)
		username = "张三"
		age = 20
		sex = "男"

		fmt.Println(username, age, sex)
	*/

	/*	var (
			username = "张三"
			age      = 20
		)

		fmt.Println(username, age)*/

	// 短变量声明法  只能声明局部变量  不能声明全局变量
	//a, b, c := 12, 13, "C"
	//fmt.Println(a, b, c)

	//var username, age = getUserInfo()
	//
	//fmt.Println(username, age)  //zhangsan 10

	// 匿名变量不占用命名空间，不会分配内容，所以匿名变量之间不存在重复声明
	//var username, _ = getUserInfo()
	//fmt.Println(username)

	//const pi = 3.14159
	//fmt.Println(pi)

	//const a = iota
	//fmt.Println(a)

	const (
		n1 = iota
		n2
		n3
		n4
	)
	fmt.Println(n1, n2, n3, n4)
}
