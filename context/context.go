package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
*
select+channel 的方案，Context 方案主要有 4 个改动点。

watchDog 的 stopCh 参数换成了 ctx，类型为 context.Context。
原来的 case <-stopCh 改为 case <-ctx.Done()，用于判断是否停止。
使用 context.WithCancel(context.Background()) 函数生成一个可以取消的 Context，用于发送停止指令。这里的 context.Background() 用于生成一个空 Context，一般作为整个 Context 树的根节点。
原来的 stopCh <- true 停止指令，改为 context.WithCancel 函数返回的取消函数 stop()。
*/
func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	ctx, stop := context.WithCancel(context.Background())

	go func() {

		defer wg.Done()

		contextWatchDog(ctx, "【监控狗1】")

	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒

	stop() //发停止指令

	wg.Wait()

}

func contextWatchDog(ctx context.Context, name string) {

	//开启for select循环，一直后台监控

	for {

		select {

		case <-ctx.Done():

			fmt.Println(name, "停止指令已收到，马上停止")

			return

		default:

			fmt.Println(name, "正在监控……")

		}

		time.Sleep(1 * time.Second)

	}

}
