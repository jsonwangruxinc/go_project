package _case

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCase() {
	var count int64 = 5
	atomic.StoreInt64(&count, 10)
	fmt.Println("获取变量的值:", atomic.LoadInt64(&count))
	// count += 10
	fmt.Println("在原有基础上累加：", atomic.AddInt64(&count, 10))
	//交换变量值，并返回原有值,count=30
	fmt.Println("交换变量值，并返回原有值", atomic.SwapInt64(&count, 30))
	fmt.Println("获取变量的值:", atomic.LoadInt64(&count))
	//比较并替换原有的值，并返回替换成功，通过旧值对比，如果旧值相等则进行替换
	fmt.Println("比较并替换原有的值，并返回替换成功：", atomic.CompareAndSwapInt64(&count, 30, 40))
	fmt.Println("获取变量的值:", atomic.LoadInt64(&count))
	//比较并替换原有的值，并返回替换成功，通过旧值对比，如果旧值不相等则不替换
	fmt.Println("比较并替换原有的值，并返回替换失败：", atomic.CompareAndSwapInt64(&count, 30, 100))
	fmt.Println("获取变量的值:", atomic.LoadInt64(&count))
}

type atomicCounter struct {
	count int64
}

func (c *atomicCounter) Inc() {
	atomic.AddInt64(&c.count, 1)
}
func (c *atomicCounter) Load() int64 {
	return atomic.LoadInt64(&c.count)
}

// 通过锁的方式
func AtomicCase1() {
	var count int64
	//添加锁  确保原子性
	locker := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			locker.Lock() //lock和unlock 之间是一个原子操作
			count += 1
			locker.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// 通过定义原子方式
func AtomicCase2() {
	var count = atomicCounter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
}
func AtomicCase3() {
	list := []string{"A", "B", "C", "D"}
	//	定义一个原子值
	var atomicMp atomic.Value
	//	定义一个集合
	mp := map[string]int{}
	//	将集合存储到原子值里面
	atomicMp.Store(&mp)
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		atomicLabel:
			// 获取默认值，map指针类型转换
			m := atomicMp.Load().(*map[string]int)
			m1 := map[string]int{}
			for k, v := range *m {
				m1[k] = v
			}
			for _, item := range list {
				_, ok := m1[item]
				if !ok {
					m1[item] = 0
				}
				m1[item] += 1
			}
			swap := atomicMp.CompareAndSwap(m, &m1)
			if !swap {
				//	重新执行替换逻辑
				goto atomicLabel
			}
		}()
	}
	wg.Wait()
	fmt.Println(atomicMp.Load())
}
