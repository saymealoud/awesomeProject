package main

import (
	"fmt"
	"reflect"
)

func main() {

	p := person{Name: "飞雪无情", Age: 20}

	pv := reflect.ValueOf(p)

	//反射调用person的Print方法

	mPrint := pv.MethodByName("Print")

	args := []reflect.Value{reflect.ValueOf("登录")}

	mPrint.Call(args)

	//i:= 10
	//ip:=&i
	//  Go 语言是不允许两个指针类型进行转换的。
	//var fp *float64 = (*float64)(ip)
	//fmt.Println(fp)

}

func (p person) Print(prefix string) {

	fmt.Printf("%s:Name is %s,Age is %d\n", prefix, p.Name, p.Age)

}

type person struct {
	Name string

	Age int
}
