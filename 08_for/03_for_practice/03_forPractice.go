package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		水仙花数：三位数，，1000~999之间
			三位数的每一个数字的立方的和刚好等于该数字本身，如：153
			1*1*1 + 5*5*5 + 3*3*3

		怎么获取每一位的单个数字？比如 268
			个位：268 % 10 = 8
			百位：268 / 100 = 2
			十位：两种方法
				将268想办法变成26，然后 26 % 10 余数就是6
				将268想办法变成68，然后 68 / 10 就是6
	*/

	for i := 100; i <= 999; i++ {
		x := i / 100     // 表示百位上的数字
		y := i / 10 % 10 // 表示十位上的数字
		z := i % 10      // 表示个位上的数字

		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println("水仙花数是：", i)
		}
	}

	fmt.Println("----------")

	/* 水仙花数 转变思路
	百位：0~9
	十位：0~9
	个位： 0~9
	*/
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				n := a*100 + b*10 + c*1
				if a*a*a+b*b*b+c*c*c == n {
					fmt.Println("水仙花数为：", n)
				}
			}
		}
	}

	fmt.Println("----------")

	/* 求素数
	打印1~100内的素数(只能被1和本身整除)
	*/
	for i := 2; i <= 100; i++ {
		flag := true // 默认认为它是素数
		// 判断到根号i，也就是对i开平方即可，不需判断要到i的前一个数字
		for j := 2; j <=int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}

}
