package main

import (
	"fmt"
)

/*
	数据类型：
	1、基本数据类型：
		布尔类型：bool
			取值：true、false
		数值类型：
			整数:
				有符号整数：
					int8(-128~127)
					int16(-32768~32767)
					int32(-2147483648~2147483647)
					int64()
				无符号整数：
					uint8(0~255)
					uint16(0~65535)
					uint32(0~4294967295)
					uint64(0~18446744073709551615)
				byte: int8
				rune: uint8
			浮点数:float32, float64
			复数：complex
		字符串类型：string
	2、复合数据类型：
		array,slice,map,function,pointer,struct,interface,channel...
*/

func main() {

	b1 := true
	fmt.Println("b1:", b1)
	b2 := false
	fmt.Println("b2:", b2)

	var i1 int8 = 100
	fmt.Println("i1:", i1)
	var i2 uint8 = 200
	fmt.Println("i2:", i2)
	var i3 int = 1000
	fmt.Println("i3:", i3)

	var i4 uint8 = 100
	var i5 byte
	i5 = i4
	fmt.Println("i4:", i4, "i5:", i5)

	// 类型推断
	var i7 = 100
	fmt.Printf("%T,%d\n,", i7, i7)

	var f1 float32 = 3.14
	var f2 float64 = 4.67
	fmt.Printf("%T,%f\n", f1, f1)
	fmt.Printf("%T,%f\n", f2, f2)

	fmt.Printf("%T,%.2f\n", f1, f1)
	fmt.Printf("%T,%.3f\n", f2, f2)

	var f3 = 2.75
	fmt.Printf("%T,%f\n", f3, f3)
}
