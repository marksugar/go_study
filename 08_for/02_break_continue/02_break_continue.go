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
}
