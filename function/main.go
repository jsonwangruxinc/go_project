package main

import (
	"context"
	"fmt"
	_case "go_project/function/case"
	"os"
	"os/signal"
)

func main() {
	//调用函数
	fmt.Println(_case.Sum(1, 2))
	//将函数结果赋值给变量
	f1 := _case.Sum
	//执行函数
	fmt.Println(f1(1, 2))

	// 将函数作为输入输出实现中间件
	f2 := _case.LogMiddleware(_case.Sum)
	//再次附加中间件
	f2 = _case.LogMiddleware(f2)
	fmt.Println(f2(1, 2))

	f4 := _case.SumFunc(f1)
	fmt.Println(f4.Accumulation(1, 2, 3, 4, 5, 6))
	fmt.Println(f2.Accumulation(1, 2, 3, 4))
	fmt.Println(_case.Fib(10))
	//闭包陷阱
	_case.ClosureTrap()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	<-ctx.Done()
}
