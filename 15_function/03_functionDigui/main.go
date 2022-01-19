/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-19 09:58:29
 * @LastEditTime: 2022-01-19 10:58:21
 * @version: v1.0
 */

/*
函数递归调用
一个函数在函数体内又调用了本身，就称之为递归调用
*/

package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}

func main() {
	test(4) // 输出 n=2,n=2,n=3，会换行输出
	fmt.Println("-------------")
	test2(4) // 输出 n=2
}
