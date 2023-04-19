package main

import (
	"fmt"
	"sync"
)

func main() {

	// 扇入 扇出测试
	coms := fan_in_out_buy(100) //采购100套配件

	//三班人同时组装100部手机

	phones1 := fan_in_out_build(coms)

	phones2 := fan_in_out_build(coms)

	phones3 := fan_in_out_build(coms)

	//汇聚三个channel成一个

	phones := merge(phones1, phones2, phones3)

	packs := fan_in_out_pack(phones) //打包它们以便售卖

	//输出测试，看看效果

	for p := range packs {

		fmt.Println(p)

	}

}

//扇入函数（组件），把多个chanel中的数据发送到一个channel中

func merge(ins ...<-chan string) <-chan string {

	var wg sync.WaitGroup

	out := make(chan string)

	//把一个channel中的数据发送到out中

	p := func(in <-chan string) {

		defer wg.Done()

		for c := range in {

			out <- c

		}

	}

	wg.Add(len(ins))

	//扇入，需要启动多个goroutine用于处于多个channel中的数据

	for _, cs := range ins {

		go p(cs)

	}

	//等待所有输入的数据ins处理完，再关闭输出out

	go func() {

		wg.Wait()

		close(out)

	}()

	return out

}

// 工序1采购
func fan_in_out_buy(n int) <-chan string {

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
func fan_in_out_build(in <-chan string) <-chan string {

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

func fan_in_out_pack(in <-chan string) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for c := range in {

			out <- "打包(" + c + ")"

		}

	}()

	return out

}
