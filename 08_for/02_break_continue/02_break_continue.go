/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2021-12-16 21:38:06
 * @LastEditTime: 2022-01-16 13:35:52
 * @version: v1.0
 */
package main

import "fmt"

/*
break：彻底结束本次循环，如果有多重循环，则break默认是跳出距离最近的一个循环，也可以通过标签指明要跳出到哪个循环
continue：结束本次循环，继续下一次的循环。如果有多重循环，则continue是可以通过标签指明要跳到哪个循环

注意：多层循环嵌套的时候，break和continue结束的都是离它最近的那一层循环。
往往也就是结束最内层嵌套的那一层循环
*/

func breakDemo() {
	for i := 1; i <= 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println("i的值:", i)
	}
	fmt.Println("Game Over...")

	for j := 1; j <= 10; j++ {
		if j == 4 {
			continue
		}
		fmt.Println("i的值:", j)
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

	fmt.Println("----------")

	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break // break 默认会跳出最近的一个for循环
			}
			fmt.Println("j=", j)
		}
	}

	fmt.Println("----- label -----")

	// break后面可以指定标签，跳出标签对应的for循环
label2:
	for k := 0; k < 4; k++ {
		// label1:
		for l := 0; l < 10; l++ {
			if l == 2 {
				// break label1
				// 当 l等于2的时候，就 break 到label2，label2在最外层的for循环，所以这里就会跳出整个for循环
				break label2
			}
			fmt.Println("l=", l)
		}
	}

	fmt.Println("----------")

	// break语句出现在多层嵌套语句块中时，可以通过标签(自定义标签名字，比如：label,label1,label2,out等)来指明要跳出到哪一层语句块
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

func continueDemo() {
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			continue
	// 		}
	// 		fmt.Println("j=", j)
	// 	}
	// 	fmt.Println("----------")
	// }

	// for i := 0; i < 13; i++ {
	// 	if i == 10 {
	// 		continue
	// 	}
	// 	fmt.Println("i=", i)
	// }

	// for i := 0; i < 2; i++ {
	// 	for j := 1; j < 4; j++ {
	// 		if j == 2 {
	// 			continue
	// 		}
	// 		fmt.Println("i=", i, "j=", j)
	// 	}
	// }

here:
	for i := 0; i < 2; i++ {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i=", i, "j=", j)
		}
	}

}

func practiceDemo1() {
	// continue实现打印1~100之间的奇数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println("1~100之间的奇数:", i)
		}
	}
}

func practiceDemo2() {
	// 从键盘输入个数不确定的整数，并判断读入的正数和负数的个数，输入为0时结束程序,positive是正数,negative表示负数
	var (
		positiveCount int = 0
		negativeCount int = 0
		num           int
	)

	//	for {
	//		fmt.Println("请输入整数，可以为正数或者负数:")
	//		fmt.Scanln(&num)
	//
	//		if num == 0 {
	//			break // 终止整个for循环
	//		}
	//
	//		if num > 0 {
	//			positiveCount++
	//		} else {
	//			negativeCount++
	//		}
	//	}
	//	fmt.Printf("输入的正数共%d个,负数共%d个\n",positiveCount,negativeCount)

	for {
		fmt.Println("请输入整数，可以为正数或者负数:")
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("输入的正数共%d个,负数共%d个\n", positiveCount, negativeCount)
}

func main() {
	// breakDemo()
	// continueDemo()
	// practiceDemo1()
	practiceDemo2()
}
