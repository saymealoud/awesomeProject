package main

import (
	"fmt"
	"reflect"
)

func main() {

	p := person{Name: "飞雪无情", Age: 20}

	ppv := reflect.ValueOf(&p)

	ppv.Elem().Field(0).SetString("张三")

	fmt.Println(p)

	fmt.Println(ppv.Kind())

	pv := reflect.ValueOf(p)

	fmt.Println(pv.Kind())

}

type person struct {
	Name string

	Age int
}
