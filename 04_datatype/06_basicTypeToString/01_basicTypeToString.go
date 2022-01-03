package main

import (
	"fmt"
	"strconv"
)

/*
官方文档：https://studygolang.com/pkgdoc

基本数据类型转string类型的两种方式

方法一：
	fmt.Sprintf("%参数", 表达式)
	参数需要和表达式的数据类型匹配，fmt.Sprintf会返回转换后的字符串

方法二：
	strconv包中的函数

*/
func main() {

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	// 使用方法一进行转换

	// 将num1转成string 然后赋值给str变量。转换的是num1变量的值99，将int类型的99变成字符串99
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	// 方法二：strconv包中的函数方式来进行转换

	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10) // 10表示十进制
	fmt.Printf("str type %T str=%q\n", str, str)

	// f代表格式，10 表示小数点后面保留10位，不足10位的，后面用0补全10位，64代表这个小数是float64类型
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	// Itoa 函数 直接将int类型转成字符串
	var num5 int64 = 678
	str = strconv.Itoa(int(num5)) // 将num5变量的值转成字符串类型，num5本身还是int类型
	fmt.Printf("str type %T str=%q\n", str, str)

}
