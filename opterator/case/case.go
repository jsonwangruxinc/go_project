package _case

import "fmt"

/*

运算符优先级:
1.一元运算符优先级最高,例如：++ -- ! ^
2.二元运算符优先级从高到低：
*	/	%	<<	>>	&	&^
+	-	|	^
==	!=	<	<=	>	>=
&&
||

*/

// 算数运算符
func ArithmeticCase() {
	//	定义两个变量
	var a = 20
	var b = 40
	var c int

	c = a + b
	fmt.Printf("a + b 的值为 %d \n", c)
	c = a - b
	fmt.Printf("a - b 的值为 %d \n", c)
	c = a * b
	fmt.Printf("a * b 的值为 %d \n", c)
	c = a / b
	fmt.Printf("a / b 的值为 %d \n", c)
	//	抹除 取余数
	c = a % b
	fmt.Printf("a %% b 的值为 %d \n", c)
	a++
	fmt.Printf("a ++ 的值为 %d \n", c)
	a--
	fmt.Printf("a -- 的值为 %d \n", c)
}

// 关系运算符
func RelationCase() {
	var a = 21
	var b = 20
	fmt.Println("a == b", a == b)
	fmt.Println("a < b", a < b)
	fmt.Println("a>b", a > b)
	fmt.Println("a <= b", a <= b)
	fmt.Println("a >=b", a >= b)
	fmt.Println("a !=b", a != b)
}

// 逻辑运算符

/*
首先，声明两个布尔变量 a 和 b 并分别初始化为 true 和 false。

&& 运算符执行逻辑 AND 运算。表达式 a && b 的计算结果为 false，因为 a 和 b 都需要为 true，AND 条件才为 true。

|| 运算符执行逻辑或运算。表达式 || b 计算结果为 true，因为对于 OR 条件为 true，a 或 b 可以为 true。由于 a 为真，因此满足 OR 条件。

这 ！运算符执行逻辑 NOT 运算。!(a && b) 计算结果为 true，因为 false 表达式 (a && b) 的 NOT 为 true。

总结来说：

a && b 为假，因为两边都不为真
一个 || b 为真，因为至少有一侧为真
!(a && b) 为 true，因为它反转了 (a && b) 的假值
*/
func LogicCase() {
	var a = true
	var b = false
	fmt.Println("a && b", a && b)
	fmt.Println("a || b", a || b)
	fmt.Println("!(a&&b)", !(a && b))
}

// 位运算
func BitCase() {
	var a uint8 = 60
	var b uint8 = 13
	var c uint8 = 0
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()
	c = a & b
	fmt.Println("a & b:")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()
	c = a | b
	fmt.Println("a | b:")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()
	c = a ^ b
	fmt.Println("a ^ b:")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()
	//	移位
	c = a << 2
	//	a 向左移动两位
	fmt.Println("a << 2 :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", c)
	fmt.Println()
	c = a >> 2
	//	a 向右移动两位
	fmt.Println("a >> 2 :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", c)
	fmt.Println()
	//	取反
	c = ^a
	fmt.Println("取反a ^值 :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", c)
	fmt.Println()
	//取反and not 先将b去反
	c = a &^ b
	fmt.Println("取反b ^值 :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()
	c = a & ^b
	fmt.Println("取反b ^值 :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()
}

// 赋值运算
func AssignmentCase() {
	var a = 21
	var c int
	c = a
	fmt.Println("c = a 赋值给,c值为:", c)
	c += a
	fmt.Println("c += a 赋值给,c值为:", c)
	c -= a
	fmt.Println("c -= a 赋值给,c值为:", c)
	c *= a
	fmt.Println("c *= a 赋值给,c值为:", c)
	c /= a
	fmt.Println("c /= a 赋值给,c值为:", c)
	c %= a
	fmt.Println("c %= a 赋值给,c值为:", c)
	var b uint8 = 60
	fmt.Printf("b 值为 %d,二进制表示:%08b\n", b, b)
	//b 左移等于
	b <<= 2
	fmt.Printf("b <<= 2 值为 %d,二进制表示:%08b\n", b, b)
	//b 右移等于
	b >>= 2
	fmt.Printf("b >>= 2 值为 %d,二进制表示:%08b\n", b, b)
	// b 且等于
	b &= 2
	fmt.Printf("b &= 2 值为 %d,二进制表示:%08b\n", b, b)
	// b 或等于
	b |= 2
	fmt.Printf("b |= 2 值为 %d,二进制表示:%08b\n", b, b)
	// b 或等于
	b ^= 2
	fmt.Printf("b ^= 2 值为 %d,二进制表示:%08b\n", b, b)
}
