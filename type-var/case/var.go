package _case

import "fmt"

func VarDeclareCase() {
	//通过var关键词声明变量
	var i, z, x int = 1, 2, 3
	//通过var关键词声明变量，并赋值
	var j int = 100
	//通过var关键词声明变量，并赋值
	var k float32 = 100.123
	//通过:= 以推断的方式定义变量并赋值，此方式只能用于局部变量的定义(bool类型)
	b := true
	//数组
	var arr = [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{2, 3, 4, 5, 6, 7, 8}
	var arr2 [5]int
	arr2[2] = 4
	arr2[3] = 5
	fmt.Println(i, z, x, j, k, b, arr, arr1, arr2)
	//指针类型,用于表示变量地址的类型
	var intPro *int
	var floatPro *float64
	var i1 = 100
	intpro(&i1)
	fmt.Println(i1, intPro, floatPro)
	//接口类型 interface 可以接受任何类型
	var inter interface{}
	inter = i1
	fmt.Println(inter)
}

func intpro(i *int) {
	*i = *i + 1
}
