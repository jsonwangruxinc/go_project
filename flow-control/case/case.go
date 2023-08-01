package _case

import (
	"fmt"
)

func FlowControlCase() {
	ifElseCase(0)
	ifElseCase(1)
	ifElseCase(2)
	forCase()
	SwitchCase("A", "1")
	SwitchCase("C", "")
	goToCase()

}

func ifElseCase(a int) {
	if a == 0 {
		fmt.Println("执行 if语句...")
	} else if a == 1 {
		fmt.Println("执行 else if 语句...")
	} else {
		fmt.Println("执行else 语句...")
	}
}

func forCase() {
	//for 循环i等于1,i小于10,i++ 判断i等于截取，跳过
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("位置1 执行for循环语句i:", i)
	}
	// i为int类型 定义b为true 当b为true时开始循环，当i大于等于10时，b为false，退出循环
	var i int
	var b = true
	for b {
		fmt.Println("位置2 执行for循环语句i:", i)
		i++
		if i >= 10 {
			b = false
		}
	}
	//遍历一个切片 第一个是索引值，第二个是数据
	//通过for循环拿到第一个index值索引，第二通过data拿到数据,通过range迭代list，然后打印
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, data := range list {
		fmt.Printf("位置3 执行for循环语句 index: %d,data:%d \n", index, data)
	}
	//遍历集合
	mp := map[string]string{"A": "a", "B": "b", "C": "c", "D": "d"}
	for key, value := range mp {
		fmt.Printf("位置4 执行for循环语句 key: %v,value:%v \n", key, value)
	}
	//	定义字符串
	str := "今天天气很好"
	for index, data := range str {
		fmt.Printf("位置5 执行for循环语句 index:%v,data:%s \n", index, string(data))
	}
	var j int
	for {
		fmt.Println("位置6 执行for语句 j:", j)
		j++
		if j >= 10 {
			break
		}
	}
	//嵌套循环
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("位置7 执行for语句i:%d,j:%d \n", i, j)
		}
	}
	//标签方式 跳转执行
lab1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("位置8 执行for语句i:%d,j:%d \n", i, j)
			break lab1
		}
	}
}

func SwitchCase(alpha string, in interface{}) {
	switch alpha {
	case "A":
		fmt.Println("Switch语句执行A，语句")
	case "B":
		fmt.Println("Switch语句执行B，语句")
	case "C", "D":
		fmt.Println("Switch语句执行CD，语句")
		fallthrough
	case "E":
		fmt.Println("Switch语句执行E，语句")
	}
	switch in.(type) {
	case string:
		fmt.Println("in 输入为字符串")
	case int:
		fmt.Println("in 输入为int类型")
	default:
		fmt.Println("未能识别in的类型，走默认")

	}
}

func goToCase() {
	var a = 0
lab1:
	fmt.Println("执行goto语句位置1")
	for i := 0; i < 10; i++ {
		a += i
		if a == 0 {
			a += 1
			goto lab1
		}
	}
}
