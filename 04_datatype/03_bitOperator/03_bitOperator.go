package main

import "fmt"

// 一元运算符

/*
左位移和右位移总结：
左移的运算规则是左移 N 位，就是乘以 2 的 N 次方。
右移的运算规则是右移 N 位，就是除以 2 的 N 次方。
*/
func main() {

	// 按位补足 ^

	/*
		位左移 <<
			用法： bitP << n
			bitP的位向左移动n位，右侧空白部分使用0填充；如果n等于2，则结果是2的n次方
	*/
	type ByteSize float64
	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)

	fmt.Println("----------")

	type BitFlag int
	const (
		Active BitFlag = 1 << iota  // iota此时为0,则 表达式为：1 << 0,即1乘以2的0次方，结果还是1
		Send						// iota此时为1，则 表达式为：1 << 1,即1乘以2的1次方，结果为2
		Receive						// iota此时为2，则 表达式为：1 << 2,即1乘以2的2次方，结果为4
	)
	flag := Active | Send
	rec := Receive
	fmt.Printf("flag的值:%d, rec的值:%d\n", flag, rec)

	var a =111
	var b = a << 3
	fmt.Println("b的值:",b)

	/*
		位右移 >>
			用法： bitP >> n
			bitP的位向右移动n位，左侧空白部分使用0填充；如果n等于2，则结果是当前值除以2的n次方

	*/

	var c = 111
	var d = c >> 3
	fmt.Println("d的值为:",d)

}
