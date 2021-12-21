package main

import (
	"fmt"
)

func main() {
	/*
		程序在内存中存储它的值，每个内存块（或字）有一个地址，通常用十六进制数表示，如：0x6b0820 或 0xf84001d7f0
		1、Go 语言的取内存地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。内存地址随着程序每一次运行而变化。
		这个地址可以存储在一个叫做指针类型的特殊数据类型中。一个指针变量可以指向任何一个值的内存地址。
		2、根据内存地址取值的符号是 *，* 号就代表获取指针所指向的内容，这里的*号其实是一个类型更改器。

		使用一个指针引用一个值被称为间接引用。

		指针的格式化标识符为：%p，一个指针变量通常缩写为 ptr

		注意：
			1、当一个指针被定义后没有分配到任何变量时，他的值为nil。
			2、符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；
				这被称为反引用（或者内容或者间接引用）操作符；另一种说法是指针转移。

		对于任何一个变量 var， 如下表达式都是正确的：var == *(&var)
		常量无法获取到内存地址
	*/

	// int类型的指针例子
	i1 := 5
	fmt.Printf("An integer:%d,it's location in memory:%p\n", i1, &i1)

	// 根据i1的内存地址取值
	var intP *int = &i1
	fmt.Printf("根据内存地址 %p 拿到的值为: %d\n", intP, *intP)

	// string类型的指针例子
	fmt.Println("----- string类型的指针例子 -----")

	s := "good bye"
	var p *string = &s // 变量p是一个指针类型的字符串，指向了变量s的内存地址
	fmt.Println("此时的变量p的指针地址为:", p)
	fmt.Println("此时的变量p的指针地址取出来的值为:", *p)

	// 通过对 *p 赋另一个值来更改“对象”，这样 s 也会随之更改。
	*p = "ciao"                                  // 变量p被重新赋值
	fmt.Printf("变量p被重新赋值后的内存地址: %p\n", p)        // 打印变量p的内存地址
	fmt.Printf("变量p被重新赋值后，根据内存地址取到的值: %s\n", *p) // 打印变量p的根据内存地址取到的值
	fmt.Printf("变量s的值: %s\n", s)                 // 打印变量s的初始值

	/*
	new和make函数来开辟内存空间
	make和new的区别：
		1、make和new都是用于申请内存（内存分配）的
		2、new比较少用，一般用于给基本数据类型的内存申请， string 、 int ，返回的是对应类型的指针(*string、*int)
		3、make用于给 slice 、 map 、 chan 申请内存，make函数返回的是对应的这三个类型本身
	 */

	// 以下例子会提示错误的内存地址或者空指针
	//var a *int  // 声明一个int类型的指针变量a。没有初始化a，则这里是一个nil的空指针。即内存地址为空
	//*a = 100    // 根据内存地址找值。因为上面的a变量是空指针，没有内存地址。所以这里根据内存地址找值，进行赋值100就出现panic了
	//fmt.Println(*a)

	// 用new函数来开辟内存空间
	//var a1 *int
	//fmt.Println(a1)  // panic 提示空指针，因为没有被初始化开辟内存空间

	var a2 *int
	a2 = new(int)    // 用new函数开辟内存空间
	fmt.Println(a2)  // 输出a2的内存地址
	fmt.Println(*a2) // 声明变量只是用new函数开辟了空间但是没有赋值，因此值为0

	*a2 = 10
	fmt.Println("a2内存地址赋值后的值:", *a2)

	/*
	make也用于分配内存地址，和new相比，它只用于slice、map、channel的内存创建。
	由于这三种类型就是引用类型，因此没必要返回他们的指针
	 */
	var b map[string]int
	b = make(map[string]int, 10)	// 用make函数进行初始化
	b["小姜"] = 100
	b["小李"] = 200
	fmt.Println(b)	// map[小姜:100 小李:200]

	// 对空指针进行反向引用不合法，会让程序崩溃
	//var np *int = nil
	//*p = 0
	// in Windows: stops only with: <exit code="-1073741819" msg="process crashed"/>
// runtime error: invalid memory address or nil pointer dereference

}