package _case

import "fmt"

// 声明AnimalI接口
//
//	定义AnimalI 行为
type AnimalI interface {
	//	吃
	Each()
	//	喝
	Drink()
	//	睡觉
	Sleep()
	//	奔跑
	Run()
}

type animal struct {
}

func (a animal) Each() {
	fmt.Println("Animal Each默认接口实现")
}
func (a animal) Drink() {
	fmt.Println("Animal Drink默认接口实现")
}
func (a animal) Sleep() {
	fmt.Println("Animal Sleep默认接口实现")
}
func (a animal) Run() {
	fmt.Println("Animal Run默认接口实现")
}
