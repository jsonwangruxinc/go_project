package _case

import "fmt"

func NewCase() {
	//	可以通过new函数,可以创建任意类型，并返回一个指针
	mpPtr := new(map[string]*user)
	//if *mpPtr == nil {
	//	fmt.Println("map 为空")
	//}
	//(*mpPtr)["A"] = &user{}
	slicePtr := new([]user)
	if *slicePtr == nil {
		fmt.Println("切片值为nil", *slicePtr)
	}
	*slicePtr = append(*slicePtr, user{
		Name: "wang",
		Age:  30,
	})
	userPtr := new(user)
	strPtr := new(string)
	fmt.Println(mpPtr, slicePtr, userPtr, strPtr)
}

// make 仅用于,切片, 集合 ,通道的初始化
func MakeCase() {
	//初始化切片，并设置长度和容量
	slice := make([]int, 10, 20)
	slice[0] = 10
	//	初始化集合，并设置集合的初始化大小
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	//	初始化通道，设置通道的读写方向和缓冲大小
	ch := make(chan int, 10)
	ch1 := make(chan<- int, 10)
	ch2 := make(<-chan int, 10)
	fmt.Println(slice, mp, ch, ch1, ch2)
}
func SliceAndMapCase() {
	//定义切片
	var slice []int
	slice = []int{1, 2, 3, 4, 5, 6, 7}
	slice1 := make([]int, 10)
	slice1[1] = 10
	fmt.Println(slice, slice1)
	//	切片截取
	slice2 := make([]int, 6, 12)
	fmt.Println(len(slice2), cap(slice2))
	slice2[0] = 0
	slice2[1] = 1
	slice2[2] = 2
	slice2[3] = 3
	slice2[4] = 4
	slice2[5] = 5
	//切片截取
	slice3 := slice2[2:6]
	fmt.Println(len(slice3), cap(slice3), slice3)
	//	切片的附加
	slice3 = append(slice3, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(len(slice3), cap(slice3), slice3)
	//	集合，无序
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	mp["B"] = "b"
	mp["C"] = "c"
	mp["D"] = "d"
	mp["E"] = "e"
	fmt.Println(mp)
	//	无序
	for k, v := range mp {
		fmt.Println(k, v)
	}
	//	删除集合元素
	delete(mp, "C")
	fmt.Println(mp)
}
