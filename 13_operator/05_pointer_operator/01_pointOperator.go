package main

import "fmt"

func main() {
	/*
		& 取内存地址，* 从内存地址中取值
		&a 变量的内存地址
		*a 是一个指针变量，通过内存地址取值一般会定义指针类型的变量 ptr
		go不支持三元运算，只能用原生的if else结构来做。go设计理念：一种事情有且只有一种解决方式
	*/
	a := 100
	fmt.Println("a的内存地址=", &a)

	var ptr *int = &a
	fmt.Println("ptr 指向的值=", *ptr)

	// go中需要用三元运算符的例子
	//var (
	//	n int
	//	i int = 10
	//	j int = 12
	//)
	//
	//if i > j {
	//	n = i
	//} else {
	//	n = j
	//}
	//fmt.Println("n=", n)

	// 例子1：求两个数的最大值
	var (
		n1  int = 10
		n2  int = 40
		max int
	)
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("三个数中最大值:", max)

	/*
		例子2：求三个数的最大值
		思路：先求出两个数的最大值，然后让最大值和第三个数比较，再求出最大值即可
	*/
	var n3 int = 45
	if n3 > max {
		max = n3
	}
	fmt.Println("三个数中最大值:", n3)

}
