package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i的值：%d\n", i)
	}

	fmt.Println("----------")

	j := 1       // 循环变量初始化
	for j <= 5 { // 循环条件
		fmt.Println(j)
		j++ // 循环变量迭代
	}

	fmt.Println("----------")

	// 死循环，通常需要配合break使用，等价于 for ;; {}
	k := 1
	for {
		if k <= 10 {
			fmt.Println("k is ok...")
		} else {
			break
		}
		k++
	}
	// for {
	// 	fmt.Println("k--->: ", k)
	// 	k++
	// 	if k == 10 {
	// 		break
	// 	}
	// }

	fmt.Println("----------")

	// 1、打印28~23的数字
	for m := 28; m >= 23; m-- {
		fmt.Println("m的值：", m)
	}

	fmt.Println("----------")

	// 2、求1~100的和
	sum := 0
	for n := 0; n <= 100; n++ {
		sum += n
	}
	fmt.Println("1~100的和：", sum)

	fmt.Println("----------")

	// 3、打印1~100以内能被3整除的数但不能被5整除的数，统计打印的数字的个数，每行打印5个
	count := 0 // 打印的数字个数总和
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 != 0 {
			fmt.Print(num, "\t") // fmt.Print打印后不换行
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Println("count个数：", count)

	fmt.Println("----------")

	// 循环嵌套
	// 1、每行打印5个*号，一共打印5行

	for i := 0; i < 5; i++ { // 变量i控制行数
		for xin := 0; xin < 5; xin++ { // 变量 xin 控制*的数量
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("----------")

	// 2、打印九九乘法表
	// 第一行
	for j := 1; j <= 1; j++ {
		fmt.Printf("%d * %d = %d\t", j, 1, 1*j)
	}
	fmt.Println()

	// 第二行
	for j := 1; j <= 2; j++ {
		fmt.Printf("%d * %d = %d\t", j, 2, 2*j)
	}
	fmt.Println()

	// 第三行
	for j := 1; j <= 3; j++ {
		fmt.Printf("%d * %d = %d\t", j, 3, 3*j)
	}
	fmt.Println()

	fmt.Println("----------")

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

}
