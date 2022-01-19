/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-18 10:51:48
 * @LastEditTime: 2022-01-19 09:52:56
 * @version: v1.0
 */

/*
Go函数基本语法：

func 函数名 (形参列表) (返回值类型列表) {
	代码块(语句块)
	return 返回值列表
}

Go函数支持返回多个值
1、如果返回多个值时，在接收时希望忽略某个返回值，则可以用 _ 符号(匿名函数) 表示占位忽略
2、如果返回值只有一个，则返回值类型列表 可以不用()包裹。返回值有多个，用括号包裹，并用逗号分隔
*/

package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

// 括号外的int(第三个int)是返回的值类型列表，单个返回值类型列表可以不写括号，也可以写上括号
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum() sum=", sum)

	// 当函数有return语句时，就是将结果返回给调用者
	// 即 谁需要调我,这个结果就返回给谁
	return sum
}

// 函数计算两个数的和与差，并返回结果
// 如果n1,n2都是int类型，则可以写成如下形式
func getSumAndSub(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1=", n1)

	// 在这里调用getSum()函数，将结果赋值给sum
	// 那么就会将getSum()函数return出来的结果返回，赋值给sum变量
	sum := getSum(10, 20)
	fmt.Println("main() sum=", sum)

	// 调用getSumAndSub()函数
	res1, res2 := getSumAndSub(5, 3)
	fmt.Printf("res1=%d,res2=%d\n", res1, res2)

	// 返回多个值时，在接收时希望忽略某个返回值，则可以用 _ 符号(匿名函数) 表示占位忽略
	_, res3 := getSumAndSub(3, 9)
	fmt.Printf("res3=%d\n", res3)
}
