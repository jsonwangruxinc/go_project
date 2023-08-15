package _case

import (
	"fmt"
	"sync"
	"time"
)

func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})
	go ReadList(&list, cond)
	go ReadList(&list, cond)
	go ReadList(&list, cond)
	time.Sleep(time.Second)
	initList(&list, cond)
}
func initList(list *[]int, c *sync.Cond) {
	//主叫方，可以持有锁，也可以不持有锁
	c.L.Lock()
	defer c.L.Unlock()
	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	//唤醒所有条件等待的协程
	c.Broadcast()
}
func ReadList(list *[]int, c *sync.Cond) {
	// 被叫方，必须持锁
	c.L.Lock()
	defer c.L.Unlock()
	for len(*list) == 0 {
		fmt.Println("readlsit wait")
		c.Wait()
	}
	fmt.Println("list数据为：", *list)
}
