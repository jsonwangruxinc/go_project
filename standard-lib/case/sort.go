package _case

import (
	"fmt"
	"sort"
)

type sortUser struct {
	ID   int64
	Name string
	Age  uint8
}

type ByID []sortUser

// 获取长度
func (a ByID) Len() int {
	return len(a)
}

// 交换位置
func (a ByID) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// i是否小于j,返回i位置的ID是否小于j位置的ID
func (a ByID) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}
func SortCase() {
	list := []sortUser{
		{ID: 10, Name: "wang", Age: 28},
		{ID: 11, Name: "wang", Age: 30},
		{ID: 02, Name: "wang", Age: 27},
		{ID: 03, Name: "wang", Age: 12},
		{ID: 04, Name: "wang", Age: 23},
		{ID: 12, Name: "wang", Age: 50},
		{ID: 14, Name: "wang", Age: 13},
		{ID: 15, Name: "wang", Age: 18},
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Age < list[j].Age
	})
	fmt.Println(list)

	//实现sort interface接口
	sort.Sort(ByID(list))
	fmt.Println(list)
}
