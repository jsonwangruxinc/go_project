package _case

import "fmt"

type user struct {
	Name string
	Age  uint
	Addr Address
}
type Address struct {
	Province string
	City     string
}

func StructCase() {
	//值类型
	u := user{
		Name: "wang",
		Age:  27,
	}
	saeck(u)
	fmt.Println(u)
	//	指针类型
	u1 := &user{
		Name: "ruxin",
		Age:  19,
	}
	fmt.Println(u1)
	//	指针类型1 new函数可以任何类型
	u2 := new(user)
	//u2.Name = "wnag1"
	//u2.Age = 20
	fmt.Println(u2, u2.Addr.Province)
	//	结构体为值类型，定义变量后将被初始化
	var u3 user
	fmt.Println(u3)
}
func saeck(u user) {
	u.Age = 38
}
