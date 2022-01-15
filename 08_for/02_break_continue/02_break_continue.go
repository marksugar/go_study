/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2021-12-16 21:38:06
 * @LastEditTime: 2022-01-15 14:44:56
 * @version: v1.0
 */
package main

import "fmt"

/*
break：彻底结束本次循环
continue：结束本次循环，继续下一次的循环

注意：多层循环嵌套的时候，break和continue结束的都是离它最近的那一层循环。
往往也就是结束最内层嵌套的那一层循环
*/
func main() {
	for i := 1; i <= 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println("i的值：", i)
	}
	fmt.Println("Game Over...")

	for j := 1; j <= 10; j++ {
		if j == 4 {
			continue
		}
		fmt.Println("i的值：", j)
	}
	fmt.Println("Game Over...")

	// 嵌套循环，这里结束的是内层循环
	//for k := 1; k <= 5; k++ {
	//	for l := 1; l <= 5; l++ {
	//		if l == 2 {
	//			//break
	//			continue
	//		}
	//		fmt.Printf("k-->%d,l==>%d\n", k, l)
	//	}
	//}
	//fmt.Println("Game Over...")

	// break语句出现在多层嵌套语句亏啊中时，可以通过标签来指明要跳出到哪一层语句块
	// 嵌套循环，因为由out，所以结束的是最外层的循环
out:
	for k := 1; k <= 5; k++ {
		for l := 1; l <= 5; l++ {
			if l == 2 {
				break out
				//continue out
			}
			fmt.Printf("k-->%d,l==>%d\n", k, l)
		}
	}
	fmt.Println("Game Over...")

	fmt.Println("----------")

	// Go语言中没有while, do while循环。但是可以通过for循环来实现它的效果

	// while 方式实现
	var m int = 1
	for {
		if m > 10 {
			break
		}
		fmt.Printf("m=%d\n", m)
		m++
	}

	/*
		for {
			循环操作(语句)
			循环变量迭代
			if 循环条件表达式 {
				break  // 跳出for循环
			}
		}
		以上是限制性在判断，因此至少会执行一次
		当循环条件成立，则执行break，会跳出for循环(结束循环)
	*/

	// do...while实现
	var n int = 1
	for {
		fmt.Println("hello,ok", n)
		n++
		if n > 10 {
			break
		}
	}
}

