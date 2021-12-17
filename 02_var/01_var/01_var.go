package main

import (
	"fmt"
   "runtime"
	"os"
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

	var name string = "allenjol"
	fmt.Println("name is: ", name)
	
	// 类型推断
	var apiVer = "v1.0"
	fmt.Println("apiversion: ", apiVer)

	// 简短定义，只能在函数体中使用简短定义
	isBoy := true
	fmt.Println("isBoy: ", isBoy)

	
	/*
	变量的零值,也就是默认值
	当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil
	*/
	
	var a int
	fmt.Println("a:", a)
	
	var b float64
	fmt.Println("b:", b)
	
	var c string
	fmt.Println("c:", c)
	

	// 批量定义变量
	var (
		gender = "women"
		hobby  = "reading"
		length = 26
	)

	fmt.Println("gender: ", gender, "hooby: ", hobby, "length: ", length)
	
	// 变量的类型也可以在运行时实现自动推断,这种写法主要用于声明包级别的全局变量
	var (
		HOME = os.Getenv("HOME")
		USER = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
	
	// 多个变量定义在一行
	var t, y, u int = 1, 2, 3
	fmt.Println(t, y, u)
	
	/*
	通过runtime包在运行时获取所在的操作系统类型，以及如何通过 os 包中的函数 os.Getenv() 来获取环境变量中的值，并保存到 string 类型的局部变量 path 中
	windows下输出：The operating system is: windows
	linux下输出：The operating system is: linux
	*/
	
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	
	// 交换变量值。注意两个变量需要类型相同才能交换。
	var j, k = 1, 2
	fmt.Println("j: ", j, "k: ", k)
	j, k = k, j
	fmt.Println("j: ", j, "k: ", k)

	s1, s2 := "allen", "jol"
	fmt.Println("s1: ", s1, "s2: ", s2)
	s1, s2 = s2, s1
	fmt.Println("s1: ", s1, "s2: ", s2)

	// 变量内存分析
	var num int  // 定义变量

	num = 100    // 用 = 号给变量赋值
	fmt.Println("num: ", num)
	fmt.Printf("num的数值是: %d,内存地址是:%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是: %d,内存地址是:%p\n", num, &num)
}
