package _case

import "fmt"

// 协程间通信
func Communication() {
	//定义一个可读可写的通道
	ch := make(chan int, 0)
	go communicationF1(ch)
	go communicationF2(ch)
}

// F1接受一个只写通道
func communicationF1(ch chan<- int) {
	//	通过循环向通道写入0~99
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

// F2接受一个只写通道
func communicationF2(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// 并发场景下的同步机制
func ConcurrentSync() {
	//	带缓冲的通道
	ch := make(chan int, 10)
	//	向ch 写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	// 向ch 写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	//	从ch读取消息并打印
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}
