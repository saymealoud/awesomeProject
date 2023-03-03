package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func sum(a, b int) (int, error) {

	if a < 0 || b < 0 {

		return 0, errors.New("a或者b不能是负数")

	}

	return a + b, nil

}

func main() {

	//fmt.Println("a")
	//fmt.Println("b")
	//fmt.Println("c")
	//
	//fmt.Print("A", "B", "C")
	//fmt.Println("A", "B", "C")

	//var a = "aaa" // go语言中变量定义以后必须要使用
	//
	//fmt.Println(a)
	//
	//fmt.Printf("%v", a)

	//var a int = 10
	//var b int = 3
	//var c int = 5

	//fmt.Println("a=", a, "b=", b, "c=", c)

	//fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
	//fmt.Printf("a=%v b=%v c=%v\n", a, b, c)

	//类型推导方式定义变量
	//a := 10
	//b := 20
	//c := 30
	//fmt.Printf("a=%v b=%v c=%v\n", a, b, c)

	//使用printf 打印一个变量的类型

	//fmt.Printf("a=%v a的类型是%T", a, a)
	array := [5]string{"a", "b", "c", "d", "e"}

	slice := array[2:5]

	fmt.Println(slice)

	slice1 := []string{"a", "b", "c", "d", "e"}

	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(slice1)

	nameAgeMap := make(map[string]int)

	nameAgeMap["飞雪无情"] = 20

	age, ok := nameAgeMap["飞雪无情"]

	if ok {

		fmt.Println(age)

	}

	nameAgeMap["飞雪无情"] = 20

	nameAgeMap["飞雪无情1"] = 21

	nameAgeMap["飞雪无情2"] = 22

	fmt.Println(len(nameAgeMap))

	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}

	s := "Hello飞雪无情"

	bs := []byte(s)

	fmt.Println(bs)

	fmt.Println(s[0], s[1], s[15])
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range s {

		fmt.Println(i, r)

	}
	fmt.Println(sum(1, 2))

	fmt.Println("飞雪", "无情")

}
