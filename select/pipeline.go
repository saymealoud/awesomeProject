package main

import (
	"fmt"
)

/**
流水线由一道道工序构成，每道工序通过 channel 把数据传递到下一个工序；
每道工序一般会对应一个函数，函数里有协程和 channel，协程一般用于处理数据并把它放入 channel 中，整个函数会返回这个 channel 以供下一道工序使用；
最终要有一个组织者（示例中的 main 函数）把这些工序串起来，这样就形成了一个完整的流水线，对于数据来说就是数据流。
*/

func main() {

	coms := buy(10) //采购10套配件

	phones := build(coms) //组装10部手机

	packs := pack(phones) //打包它们以便售卖

	//输出测试，看看效果

	for p := range packs {

		fmt.Println(p)

	}

}

// 工序1采购
func buy(n int) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for i := 1; i <= n; i++ {

			out <- fmt.Sprint("配件", i)

		}

	}()

	return out

}

// 工序2组装
func build(in <-chan string) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for c := range in {

			out <- "组装(" + c + ")"

		}

	}()

	return out

}

//工序3打包

func pack(in <-chan string) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for c := range in {

			out <- "打包(" + c + ")"

		}

	}()

	return out

}
