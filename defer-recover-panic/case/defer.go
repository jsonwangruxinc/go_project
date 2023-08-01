package _case

import (
	"fmt"
	"io"
	"log"
	"os"
)

// defer 关键词用来声明一个延迟调用函数
// 该函数可以是匿名函数，也可以是具名函数
// defer延迟函数的执行顺序为后进先出
func DeferCase1() {
	fmt.Println("开始执行DeferCase1")
	defer func() {
		fmt.Println("调用了匿名函数1")
	}()
	defer f1()
	defer func() {
		fmt.Println("调用了匿名函数2")
	}()
	fmt.Println("DeferCase1执行结束")

}

// 参数预计算
func DeferCase2() {
	i := 1
	//传参方式
	defer func(j int) {
		fmt.Println("打印defer j", j)
	}(i + 1)
	//闭包的方式
	defer func() {
		fmt.Println("打印defer j", i)
	}()
	i = 99
	fmt.Println("打印i值为:", i)
}

// 返回值
// defer 函数在执行return之后
func DeferCase3() {
	i, j := f2()
	fmt.Printf("i:%d,j:%d,g:%d\n", i, *j, g)
}

// 异常处理
func ExceptionCase() {
	defer func() {
		//捕获异常 恢复协程
		err := recover()
		//如果err不为空，为异常处理
		if err != nil {
			fmt.Println("异常处理 defer recover", err)
		}
	}()
	fmt.Println("开始执行ExceptionCase方法")
	panic("ExceptionCase抛出异常")
	fmt.Println("ExceptionCase方法执行结束")

}

// 资源释放
func FileReadCase() {
	file, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	//通过defer调用资源释放的方法
	defer func() {
		file.Close()
		fmt.Println("释放文件资源")
	}()
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(buf[:n])
	}
}

func f1() {
	fmt.Println("调用了具名函数f1")
}

var g = 100

func f2() (int, *int) {
	defer func() {
		g = 200
	}()
	fmt.Println("f2,g的值:", g)
	return g, &g
}
