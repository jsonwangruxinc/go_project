// 声明包
package main

//引入包
import (
	"fmt"
	"go_project/lan-structure/package1"
	_ "go_project/lan-structure/package2"
)

/*
多行注释:multiline comment
*/
// 初始化函数,goland每个包的引用会优先调用init函数
func init() {
	fmt.Println("调用 lan-structure init函数")
}

// 函数（main 函数为整个程序的入口）
func main() {
	//变量
	var i = 10
	//语句或表达式
	fmt.Println("调用 lan-structure main 方法")
	fmt.Println(fmt.Sprintf("打印参数i:%d", i))
	//包名可与包引入目录不一致（包名和文件夹命名可以不同）
	//goland 公有成员与私有成员通过成员的首字母来区分的
	//首字母大写表示公有成员,首字母小写表示私有成员,只有两种情况
	P1.F1()
}
