package main

import (
	"fmt"
)

/*
变量定义了就必须要使用，否则编译会报错
导入包不使用也会报错
*/

// 全局变量

// 声明变量时指定变量类型
var age int = 26
var sex string = "female"

func main() {

	fmt.Println("age: ", age)
	fmt.Println("sec: ", sex)

	// 类型推断
	var apiVer = "v1.0"
	fmt.Println("apiversion: ", apiVer)

	// 简短定义，只能在函数体中使用简短定义
	isBoy := true
	fmt.Println("isBoy: ", isBoy)

	/* 变量的零值,也就是默认值
	整形默认值为0
	浮点类型默认值为0，其实是0.0
	字符串默认值是空字符串 ""
	*/
	var a int
	fmt.Println("a: ", a)
	var b float64
	fmt.Println("b: ", b)
	var c string
	fmt.Println("c: ", c)

	// 批量定义变量
	var (
		gender = "women"
		hobby  = "reading"
		length = 26
	)

	fmt.Println("gender: ", gender, "hooby: ", hobby, "length: ", length)

	// 交换变量值
	var j, k = 1, 2
	fmt.Println("j: ", j, "k: ", k)
	j, k = k, j
	fmt.Println("j: ", j, "k: ", k)

	s1, s2 := "allen", "jol"
	fmt.Println("s1: ", s1, "s2: ", s2)
	s1, s2 = s2, s1
	fmt.Println("s1: ", s1, "s2: ", s2)

	// 变量内存分析
	var num int

	num = 100
	fmt.Println("num: ", num)
	fmt.Printf("num的数值是: %d,内存地址是:%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是: %d,内存地址是:%p\n", num, &num)

	var name string = "allenjol"
	fmt.Println("name is: ", name)
}
