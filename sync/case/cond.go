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

type queue struct {
	list []int
	code *sync.Cond
}

func NewQueue() *queue {
	q := &queue{
		list: []int{},
		code: sync.NewCond(&sync.Mutex{}),
	}
	return q
}
func (q *queue) Put(item int) {
	q.code.L.Lock()
	defer q.code.L.Unlock()
	q.list = append(q.list, item)
	//	当数据写入成功后，唤醒一个协程来处理数据
	q.code.Signal()
}
func (q *queue) GetMany(n int) []int {
	q.code.L.Lock()
	defer q.code.L.Unlock()
	for len(q.list) < n {
		q.code.Wait()
	}
	list := q.list[:n]
	q.list = q.list[n:]
	return list
}
func CodeQueueCase() {
	//	初始化一个队列
	q := NewQueue()
	//	声明一个WaitGroup
	var wg sync.WaitGroup
	for n := 1; n <= 10; n++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			lsit := q.GetMany(n)
			fmt.Printf("%d:%d \n", n, lsit)
		}(n)
	}
	for i := 0; i < 100; i++ {
		q.Put(i)
	}
	wg.Wait()
}
