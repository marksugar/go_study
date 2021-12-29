package main

import "fmt"

/*
基本数据类型转换.go中需要显示的类型转换
表达式T(v)，将变量v转换为类型T
T 就是数据类型，如：int32,int64,float64等
v 就是需要转换的变量

注意：被转换的是 变量存储的数据(也就是变量的值)，变量本身的数据类型并没发生变化
*/
func basicTransfer() {

	var i int32 = 100
	// 将变量i转成float类型
	var n1 float32 = float32(i) // 将i的值转换成float32类型，然后赋值给n1
	var n2 int8 = int8(i)
	var n3 int64 = int64(i) // 低精度(int32)转换成高精度(int64)也需要显示转换类型

	fmt.Printf("i=%v, n1=%v\n", i, n1)
	fmt.Printf("n2=%v\n", n2)
	fmt.Printf("n3=%v\n", n3)

	// 被转换的是 变量存储的数据(也就是变量的值)，变量本身的数据类型并没发生变化
	fmt.Printf("i type is: %T\n", i)

	// 如果将int64 转成int8(-128~127),编译时不会报错，但是转换的结果会按照溢出处理
	// 因此最后得到的值就不是我们想要的值。因此转换时需要考虑不同的数据类型支持的数据范围
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)
}

func pricticeOne() {
	fmt.Println("----------")

	var n1 int32 = 12
	var n2 int64
	var n3 int8

	//n2 = n1 + 20	// n1 是int32类型，n1的值12加上20，还是int32类型，然后将int32赋值给int63类型的n2，所以报错
	//n3 = n1 + 20	// n1 是int32类型，n1的值12加上10，还是int32类型，然后将int32赋值给int8类型的n3，所以报错

	// 转换一下让以上变成正确的
	n2 = int64(n1 + 20)
	n3 = int8(n1 + 10)
	fmt.Println("n2=", n2, "n3=", n3)

}

func pricticeTwo() {
	//var n1 int32 = 12
	//var n3 int8
	//var n4 int8
	//
	//n4 = int8(n1) + 127	// 12+127超过了n4的int8类型的范围，范围溢出了，所以结果不对，但能通过编译
	//n3 = int8(n1) + 128	// n3范围-128~127,所以编译都不能通过

	// 解决上面报错的问题
	fmt.Println("----------")

	var n1 int32 = 12
	var n3 int64
	var n4 int64
	n4 = int64(n1) + 127
	n3 = int64(n1) + 128

	fmt.Println("n3=",n3, "n4=", n4)
}

func main() {
	basicTransfer()
	pricticeOne()
	pricticeTwo()
}
