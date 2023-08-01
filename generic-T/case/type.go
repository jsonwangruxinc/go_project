package _case

import "fmt"

type user struct {
	ID   int64
	NAME string
	Age  uint8
}

type address struct {
	ID       int
	Province string
	City     string
}

// 集合转列表
func mapToList[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}
func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}
func TTypeCase() {
	UserMp := make(map[int64]user, 0)
	UserMp[1] = user{ID: 1, NAME: "wang", Age: 18}
	UserMp[2] = user{ID: 2, NAME: "Li", Age: 27}
	userList := mapToList[int64, user](UserMp)

	ch := make(chan user)
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}
	addrMp := make(map[int]address, 0)
	addrMp[1] = address{ID: 1, Province: "湖南", City: "长沙"}
	addrMp[2] = address{ID: 2, Province: "山东", City: "枣庄"}
	addrList := mapToList[int, address](addrMp)

	ch1 := make(chan address)
	go myPrintln(ch1)
	for _, addr := range addrList {
		ch1 <- addr
	}
}

// 泛型切片的定义
type List[T any] []T

// 泛型集合的定义
// 声明了两个泛型 k和v
type MapT[k comparable, v any] map[k]v

// 泛型通道的定义
type Chan[T any] chan T

func TTypeCase1() {
	UserMp := make(MapT[int64, user], 0)
	UserMp[1] = user{ID: 1, NAME: "wang", Age: 18}
	UserMp[2] = user{ID: 2, NAME: "Li", Age: 27}

	var userList List[user]
	userList = mapToList[int64, user](UserMp)

	ch := make(Chan[user])
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}
	addrMp := make(MapT[int, address], 0)
	addrMp[1] = address{ID: 1, Province: "湖南", City: "长沙"}
	addrMp[2] = address{ID: 2, Province: "山东", City: "枣庄"}

	var addrList List[address]
	addrList = mapToList[int, address](addrMp)

	ch1 := make(Chan[address])
	go myPrintln(ch1)
	for _, addr := range addrList {
		ch1 <- addr
	}
}
