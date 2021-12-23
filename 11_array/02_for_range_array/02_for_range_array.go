package main

import "fmt"

func main() {
	/*
		遍历数组：依次访问数组中的元素
		方法一：arr[0], arr[1],arr[2]...
		方法二：通过循环配合索引遍历
			for i :=0;i<len(arr);i++ {
		    	fmt.Print(arr[i], "\t")
		    }
		方法三：使用range
			不需要数组索引，到达数组末尾自动结束for range循环。每次都从数组中获取索引和对应的值
	*/

	arr1 := [5]int{1, 2, 3, 4, 5}

	// 逐个打印
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	fmt.Println(arr1[4])

	fmt.Println("----------")

	// for循环来打印
	for i := 0; i < len(arr1); i++ {
		fmt.Println("数组的值为：", arr1[i])
	}

	fmt.Println("----------")

	// range循环打印下标的索引和值
	for index, value := range arr1 {
		fmt.Printf("index:%d,value:%d\n", index, value)
	}

	fmt.Println("----------")
	
	// 只要数值不要索引,用匿名函数 _ 丢弃
	sum := 0
	for _, value := range arr1 {
		sum += value
	}
	fmt.Printf("数值的和是:%d\n", sum)
}
