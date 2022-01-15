/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-14 10:48:54
 * @LastEditTime: 2022-01-14 13:40:40
 * @version: v1.0
 */

package main

import (
	"fmt"
)

func cal(n1 float64, n2 float64, operator string) float64 {
	var res float64

	switch operator {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	default:
		fmt.Println("输入操作符错误!")
	}
	return res
}

func main() {
	var (
		n1 float64 = 2.5
		n2 float64 = 3.5
	)
	result := cal(n1, n2, "+")
	fmt.Printf("result=%v\n", result)

	var (
		c1 float64 = 10.6
		c2 float64 = 5.3
	)
	result2 := cal(c1, c2, "-")
	fmt.Printf("result=%v\n", result2)

}
