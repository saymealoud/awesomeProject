package main

import (
	"fmt"
	"reflect"
	"strings"
)

/**
反射的三大定律。

任何接口值 interface{} 都可以反射出反射对象，也就是 reflect.Value 和 reflect.Type，通过函数 reflect.ValueOf 和 reflect.TypeOf 获得。
反射对象也可以还原为 interface{} 变量，也就是第 1 条定律的可逆性，通过 reflect.Value 结构体的 Interface 方法获得。
要修改反射的对象，该值必须可设置，也就是可寻址，
*/

func main() {

	p := person_josn{Name: "飞雪无情", Age: 20}

	pv := reflect.ValueOf(p)

	pt := reflect.TypeOf(p)

	//自己实现的struct to json

	jsonBuilder := strings.Builder{}

	jsonBuilder.WriteString("{")

	num := pt.NumField()

	for i := 0; i < num; i++ {

		jsonTag := pt.Field(i).Tag.Get("json") //获取json tag

		jsonBuilder.WriteString("\"" + jsonTag + "\"")

		jsonBuilder.WriteString(":")

		//获取字段的值

		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))

		if i < num-1 {

			jsonBuilder.WriteString(",")

		}

	}

	jsonBuilder.WriteString("}")

	fmt.Println(jsonBuilder.String()) //打印json字符串

}

type person_josn struct {
	Name string

	Age int
}
