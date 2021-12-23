package main

import "fmt"

func main() {
	/*
		一维数组：
			a1 := [3]int{1,2,3}

		二维数组：
			a2 := [3][4]int{{},{},{}}
			该二维数组的长度就是3,存储的元素是一维数组，一维数组中的元素就是数值，每个一维数组的长度为4

		len(a2) 是二维数组的长度，长度就是二维数组当中有几个一维数组，这里有3个一维数组，那么长度为3
		a2[0] 就是二维数组中的第一个一维数组
		len(a2[0]) 就是第一个一维数组的长度，其实就是一维数组中存储的数值的个数
		a2[0][0] 就是第一个一维数组中的第一个元素
	*/

	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	fmt.Println(a2)
	fmt.Printf("二维数组的地址：%p\n", &a2)
	fmt.Printf("二维数组的长度：%d\n", len(a2))

	fmt.Printf("一维数组的长度：%d\n", len(a2[0]))
	fmt.Println(a2[0][3]) // 第一个一维数组中的第四个元素，4
	fmt.Println(a2[1][2]) // 第二个一维数组中的第三个元素，7
	fmt.Println(a2[2][1]) // 第三个一维数组中的第一个元素，10

	// 遍历二维数组
	for i := 0; i < len(a2); i++ {
		for j := 0; j < len(a2[i]); j++ {
			fmt.Print(a2[i][j], "\t")
		}
		fmt.Println()
	}

	fmt.Println("----------")

	// for range遍历
	for _, arr := range a2 {
		for _, val := range arr {
			fmt.Println(val, "\t")
		}
		fmt.Println()
	}
}
