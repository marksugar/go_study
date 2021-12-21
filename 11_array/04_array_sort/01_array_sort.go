package main

import "fmt"

func main() {
	/*
			数组排序：让数组中的元素具有一定的顺序
				arr :=[5]int{15,23,8,10,17}
					升序：[8,10,15,17,23]
					降序：[23,17,15,10,8]

		排序算法：冒泡排序、插入排序、选择排序、希尔排序、堆排序、快速排序

			冒泡排序： 依次比较两个相邻的元素，如果他们的顺序(如从小到大)，就将他们调换位置
	*/

	// 冒泡排序算法

	arr := [5]int{15, 23, 8, 18, 7}
	// 第一轮排序
	// j代表索引，索引到3即可，因为索引为4就是最后一位，最后一位不需要排序了
	//for j := 0; j < 4; j++ {
	//	if arr[j] > arr[j+1] {
	//		arr[j], arr[j+1] = arr[j+1], arr[j]
	//	}
	//}
	//fmt.Println(arr) // [15 8 18 7 23]
	//
	//// 第二轮排序
	//for j := 0; j < 3; j++ {
	//	if arr[j] > arr[j+1] {
	//		arr[j], arr[j+1] = arr[j+1], arr[j]
	//	}
	//}
	//fmt.Println(arr) // [8 15 7 18 23]

	// 上面都是重复的代码，所以这里通过在外层加一个for循环去循环排序的轮数
	// 这里的i就表示轮数
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr) // 打印出每次排序后的数组
	}
	//fmt.Println(arr) // 打印出最后一次排序后的数组
}
