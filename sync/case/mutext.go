package _case

import (
	"fmt"
	"sync"
)

func MutexCase() {
	//SingLeRoutine()
	//multipleRoutine()
	multipleSafeRoutine()
}

// 单协程操作
func SingLeRoutine() {
	mp := make(map[string]int, 0)
	list := []string{"A", "B", "C", "D"}
	for i := 0; i < 20; i++ {
		for _, item := range list {
			_, ok := mp[item]
			if !ok {
				mp[item] = 0
			}
			mp[item] += 1
		}
	}
	fmt.Println(mp)
}

// 多协程操作，非协程安全
func multipleRoutine() {
	mp := make(map[string]int, 0)
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		//开启协程
		go func() {
			defer wg.Done()
			for _, item := range list {
				_, ok := mp[item]
				if !ok {
					mp[item] = 0
				}
				mp[item] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp)
}

// 互斥锁协程安全
func multipleSafeRoutine() {
	type SafaMap struct {
		data map[string]int
		sync.Mutex
	}
	mp := SafaMap{
		data:  map[string]int{},
		Mutex: sync.Mutex{},
	}
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		//开启协程
		go func() {
			defer wg.Done()
			mp.Lock()
			defer mp.Unlock()
			for _, item := range list {
				_, ok := mp.data[item]
				if !ok {
					mp.data[item] = 0
				}
				mp.data[item] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp.data)
}
