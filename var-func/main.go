package main

import (
	"fmt"
	_case "go_project/var-func/case"
)

func main() {
	a := 10
	b := 20
	fmt.Println(_case.SumCase(a, b))
	_case.ReferenceCase(a, &b)
	fmt.Println(a, b)
	//	全局变量
	fmt.Println(_case.G)
	_case.ScopCase(a, b)
	fmt.Println(_case.G)

	user := _case.NewUser("wang", 27)
	fmt.Println(user.GetName(), user.GetAge())
}
