package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 未能实现
func main() {

	var wg sync.WaitGroup

	wg.Add(4) //记得这里要改为4，原来是3，因为要多启动一个协程

	//省略其他无关代码

	ctx := context.Background()
	valCtx := context.WithValue(ctx, "userId", 2)

	go func() {

		defer wg.Done()

		getUser(valCtx)

	}()

	//省略其他无关代码

}

func getUser(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():

			fmt.Println("【获取用户】", "协程退出")

			return

		default:

			userId := ctx.Value("userId")

			fmt.Println("【获取用户】", "用户ID为：", userId)

			time.Sleep(1 * time.Second)

		}

	}

}
