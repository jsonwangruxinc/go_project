package _case

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Printf("不使用泛型，数字比较:", getMaxNumInt(a, b))
	fmt.Printf("不使用泛型，数字比较:", getMaxNumFloat(c, d))

	fmt.Printf("使用泛型，数字比较:", getMaxNum(a, b))
	fmt.Printf("使用泛型，数字比较:", getMaxNum[float64](c, d))

}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 定义泛型
func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CusNum interface {
	// 支持 uint,int32,float64与int64以及其衍生类型
	// ~ 表示支持类型的衍生类型
	// | 表示取并集
	//多行之间去交集
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint64
}

// MayInt64 为int64的衍生类型，是具有基础类型int64的新类型，与int64是不同的类型
type MayInt64 int64

// MyInt32 为 int32 的别名，与int32是同一类型
type MyInt32 = int32
