package main

import "fmt"

func main() {
	/*
			概念：存储一组相同数据类型的数据结构
			数组定长，通过索引来读取，索引从0开始，第一个元素索引为0，第二个元素索引为1，以此类推
			数组的下标(索引)取值范围从0开始到数组的长度减1

		声明数组：
			var 数组名 [数组长度] 数据类型
			var 数组名 = [数组长度] 数据类型{预算内宿1，元素2，元素3...}
			数组名 := [...]数据类型{元素1,元素2,元素3...}

		通过索引访问：
			索引，也叫下标：index
			通过索引访问数组中的值其实就是通过数组的索引找到数组中的元素所在的内存
			数组索引默认从0开始的整数，直到长度减1。超过范围会报错索引越界
			数组名[index]

		长度和容量：Go的内置函数
			len(array/map/slice/string),长度
			cap()，容量

		int类型的数组，默认都是零值
		string类型的数组，默认都是空字符串

	*/

	var num1 int
	num1 = 100
	num1 = 200
	fmt.Println("num1:", num1)
	// 打印num1的内存地址
	fmt.Printf("num1的内存地址：%p\n", &num1)

	// 创建数组
	var arr1 [4]int // 代码格式化后 int会紧贴着括号
	// 打印数组的内存地址
	fmt.Printf("arr1数组的内存地址：%p\n", &arr1)

	// 数组访问
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4

	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	//fmt.Println(arr1[4]) //会报错，提示索引越界.out of bounds for 4-element array

	// 查看数组长度
	fmt.Println("数组的长度：", len(arr1))
	//查看数组的容量
	fmt.Println("数组的容量：", cap(arr1))

	// 根据索引更改数组中的值
	arr1[1] = 10
	fmt.Println("更改索引后arr1[1]的值：", arr1[1])

	// 数组的其他创建方式
	// 没有指定索引的值，则默认索引上的值用0填充
	var a [4]int // 等同于 var a = [4]int
	fmt.Println(a)

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	// 定义长度为5的数组，但是里面只放3个值
	var c = [5]int{1, 2, 4}
	fmt.Println(c)

	// 创建数组，指定索引位置上的值。没指定值的索引位置上值为0
	var d = [5]int{1: 1, 3: 2, 4: 15}
	fmt.Println(d)

	// 定义长度为5的数组，只有4个值，则最后一个值为空字符串
	var e = [5]string{"rose", "jack", "allen", "jol"}
	fmt.Println(e)

	var f = [5]int{'A', 'B', 'C', 'D'} // byte。一定要用单引号，双引号就是字符串了
	fmt.Println(f)                     // [65,66,67,68]

	g := [...]int{1, 2, 3, 4, 5}
	fmt.Println(g)
	fmt.Println(g[3])
	fmt.Println(len(f))

	h := [...]int{1: 4, 2: 7}
	fmt.Println(h)
	fmt.Println(len(h))
}
