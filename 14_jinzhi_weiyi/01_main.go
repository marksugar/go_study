/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-09 11:26:05
 * @LastEditTime: 2022-01-09 23:19:01
 * @version: v1.0
 */

package main

import "fmt"

func jinZhi() {
	// 二进制只有0和1。zgo中不能直接用二进制来表示一个整数
	var i int = 5
	fmt.Printf("%b\n", i)

	// 八进制,0-7,满8进1，用数字0开头来表示8进制
	var j int = 011 // 011 = 9
	fmt.Println("j=", j)

	// 十六进制，0-9及A-F,A-F不区分大小写。
	// 16进制中永远不会出现16，最多到15.满16进1，用0x或者0X开头表示
	var k int = 0x11 // 0x11 就是 16进制 的 16+1 =17
	fmt.Println("k=", k)
}

func weiYunsuan() {
	// 2的补码：00000010，3的补码：00000011,-2的原码：10000010，反码：11111101，补码：11111110
	fmt.Println(2 & 3)  // 2
	fmt.Println(2 | 3)  // 3
	fmt.Println(2 ^ 3)  // 1
	fmt.Println(-2 ^ 2) // -4
}

func weiYi() {
	/*
		右移>>：低位溢出，符号位不变，并用符号位补溢出的高位
		左移<<：符号位不变，低位补0
	*/
	a := 1 >> 2 // 1的补码为：00000001，右移2位--> 00000000 = 0
	b := 1 << 2 // 1的补码为：00000001,左移2位--> 00000100 = 4
	fmt.Println("a=", a, "b=", b)
}

func main() {
	jinZhi()
	weiYunsuan()
	weiYi()
}
