/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-18 10:51:48
 * @LastEditTime: 2022-01-18 14:16:06
 * @version: v1.0
 */

package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum() sum=", sum)

	// 当函数有return语句时，就是将结果返回给调用者
	// 即 谁需要调我这个结果就返回给谁
	return sum
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1=", n1)

	sum := getSum(10, 20)
	fmt.Println("main() sum=", sum)
}
