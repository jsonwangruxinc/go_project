package _case

import "fmt"

type Dove struct {
}

func NewDove() AnimalI {
	return &Dove{}
}
func (a *Dove) Each() {
	fmt.Println("鸽子吃虫子")
}
func (a *Dove) Drink() {
	fmt.Println("鸽子喝水")
}
func (a *Dove) Sleep() {
	fmt.Println("鸽子睡在树上")
}
func (a *Dove) Run() {
	fmt.Println("鸽子用翅膀飞")
}
