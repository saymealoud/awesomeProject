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

	//const (
	//	n1 = iota
	//	n2
	//	n3
	//	n4
	//)
	//fmt.Println(n1, n2, n3, n4)

	//var b float64 = 3.12
	//fmt.Printf("值： %v---%f,类型%T\n", b, b, b)
	//fmt.Println(unsafe.Sizeof(b))

	//var c float64 = 3.1415926123
	//fmt.Printf("%v---%f\n", c, c)
	//fmt.Printf("%v---%.2f", c, c)

	// 64 位的系统中Go语言的浮点数默认为float64
	//f1 := 3.12324523432
	//fmt.Printf("%f----%T", f1, f1)

	//var f2 float32 = 3.14e2 //表示f2 等于3.14 * 10的2次方
	//fmt.Printf("%v---%T", f2, f2)

	//m1 := 8.2
	//m2 := 3.8
	//
	//fmt.Println(m1 - m2)

	//var num1 float64 = 3.2
	//var num2 float64 = 2.1
	//
	//d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(float64(num2)))
	//fmt.Println(d1)

	// Go语言中不允许将整形转换为布尔型
	//   布尔类型也无法进行数值运算

	// 注意  转换的时候尽量从低位转换位高位
	// 其他类型转换为String 类型  fmt.Sprintf
	//  strconv   将其他类型转换为String 类型
	//var i int = 20
	//str := strconv.FormatInt(int64(i), 10)
	//fmt.Printf("值： %v 类型 %T", str, str)

	//str := "123456"
	//fmt.Printf("%v ---%T", str, str)

	//  参数一  要转换的值  参数二  进制  参数三  位数  32  64
	//num, _ := strconv.ParseInt(str, 10, 64) //strconv.ParseFloat   ParseBool  ParseUint 一样用法
	//fmt.Printf("%v ---%T", num, num)

	var a = 6
	var b = 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	var c = a * b
	fmt.Println(c)

	var d = 10
	var e = 3
	fmt.Println(d / e)

	//取余注意    余数 = 被除数 -（被除数/除数）* 除数

	// golang 中 没有 前置  ++  --

}
