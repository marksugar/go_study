/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-13 21:53:21
 * @LastEditTime: 2022-01-13 22:25:38
 * @version: v1.0
 */

package main

import "fmt"

func traverseString() {
	var str string = "allen,jol"
	for i := 0; i < len(str); i++ {
		fmt.Printf("index=%d, value=%c\n", i, str[i])
	}
}

func traverseRange() {
	str := "abcHefg"
	for index, value := range str {
		fmt.Printf("index=%d,value=%c \n", index, value)
	}
}

/*
如果字符串含有中文，那么传统的遍历字符串的方式就会出现乱码
原因是：传统对字符串的遍历是按照 字节 来进行遍历的，而一个汉字在utf-8编码
是对应 3 个字节。
解决方式：将 str 转成 []rune 切片类型即可。
*/

func traverseChiStr() {
	var str string = "allenjol!浙江"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("index=%d, value=%c\n", i, str2[i])
	}
}

// for range 是按照 字符 的方式来进行遍历的，因此如果字符串有中文也是可以正常遍历
func traverseRanStr() {
	var str string = "allenjol!杭州"
	for key, value := range str {
		fmt.Printf("index=%d, value=%c\n", key, value)
	}
}

// practice
// 1、打印 1~100 之间所有是9的倍数的整数的个数以及整数的总和
func practiceDemo1() {
	var (
		count int // 表示9的倍数的数字个数
		sum   int // 表示9的倍数的数字个数的总和
	)
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count += 1
			sum += i
		}
	}
	fmt.Printf("9的倍数的数字个数=%d,9的倍数的数字个数的总和=%d\n", count, sum)

	// var (
	// 	max   int = 100
	// 	count int = 0
	// 	sum   int = 0
	// )

	// for i := 1; i <= max; i++ {
	// 	if i%9 == 0 {
	// 		count++
	// 		sum += i
	// 	}
	// }
	// fmt.Printf("count=%v,sum =%v\n",count,sum)
}

func practiceDemo2() {
	var n int = 6
	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v\n", i, n-i, n)
	}
}

func main() {
	traverseString()
	fmt.Println("----------")
	traverseRange()
	fmt.Println("----------")
	traverseChiStr()
	fmt.Println("----------")
	traverseRanStr()
	fmt.Println("----------")

	// practice
	practiceDemo1()
	fmt.Println("----------")
	practiceDemo2()
}
