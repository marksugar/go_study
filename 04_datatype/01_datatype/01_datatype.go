package main

import (
	"fmt"
	"unsafe"
)

/*
	数据类型：
	1、基本数据类型：
		布尔类型：bool
			取值：true、false
		数值类型：1 byte = 8 bit
			整数: 有无符号区别在于他们支持的数值范围
				有符号整数：
					int8(-128~127,占1字节存储空间)
					int16(-32768~32767,即2的15次方~2的15次方减1,占2字节存储空间)
					int32(-2147483648~2147483647,即2的31次方~2的31次方减1,占4字节存储空间)
					int64(2的63次方~2的63次方减1,占8字节存储空间)
				无符号整数：
					uint8(0~255,占1字节存储空间)
					uint16(0~65535,0~2的16次方减1,占2字节存储空间)
					uint32(0~4294967295,0~2的32次方减1,占4字节存储空间)
					uint64(0~18446744073709551615,0~2的64次方减1,占8字节存储空间)
				byte: 无符号，等价于int8，范围0~255，当要存储字符时选择byte
				rune: int32，-2的31次方到2的31次方减1，表示一个Unicode编码
				
				int和uint类型，32位系统占4个字节，64位系统占8个字节
			浮点数:float32, float64
			复数：complex
		字符串类型：string
	2、复合数据类型：
		array,slice,map,function,pointer,struct,interface,channel...
*/

func intDemo() {

	fmt.Println("int类型----------")
	b1 := true
	fmt.Println("b1:", b1)
	b2 := false
	fmt.Println("b2:", b2)

	var i1 int8 = 100
	fmt.Println("i1:", i1)

	var i2 uint8 = 200
	fmt.Println("i2:", i2)

	var uit uint16 = 255
	fmt.Println("uit:", uit)

	var i3 int = 1000
	fmt.Println("i3:", i3)

	var i4 uint8 = 100
	var i5 byte
	i5 = i4
	fmt.Println("i4:", i4, "i5:", i5)

	var a int = 8900
	fmt.Println("a:", a)

	var b uint = 1
	var c byte = 255
	fmt.Println("b=", b, "c=", c)

	var n1 = 100
	// %T表示数据类型
	fmt.Printf("n1的数据类型:%T\n", n1)

	// 查看某个变量的占用字节大小和数据类型
	var n2 int64 = 10
	// unsafe.Sizeof(n2) unsafe包中的函数，可以返回n2变量占用的字节数
	fmt.Printf("n2的类型:%T,n2占用的字节数是:%d\n", n2, unsafe.Sizeof(n2))

	// 类型推断
	var i7 = 100
	fmt.Printf("%T,%d\n", i7, i7)

}

func floatDemo() {
	/*
	单精度 float32,双精度 float64
	浮点数=符号位+指数位+尾数位，浮点数都是有符号的
	float64位的精度比float32位的精度高
	Golang的浮点型默认声明为float64类型
	开发过程中推荐使用float64
	 */
	fmt.Println("float类型----------")
	var f1 float32 = 89.12
	var f2 float64 = 4.67
	fmt.Printf("%T,%f\n", f1, f1)
	fmt.Printf("%T,%f\n", f2, f2)

	fmt.Printf("%T,%.2f\n", f1, f1)
	fmt.Printf("%T,%.3f\n", f2, f2)

	var f3 = -0.00089
	fmt.Printf("%T,%f\n", f3, f3)

	var f4 = -7809656.093
	fmt.Printf("%T,%f\n", f4, f4)

	/*
	尾数部分可能丢失，造成精度损失
	float64位的精度比float32位的精度高
	 */

	var f5 float32 = -123.0000901
	var f6 float64 = -123.0000901

	fmt.Println("f5=", f5,"\t","f6=", f6)

	num1 := 5.12
	num2 := .123
	fmt.Println("num1=", num1, "num2=", num2)

	num3 := 5.1234e2	// 5.1235 * 10的2次方
	num4 := 5.1234E2
	num5 := 5.1234E-2	// 5.1235 / 10的2次方
	fmt.Println("num3=", num3, "num4=", num4, "num5=", num5)
}

func main() {
	intDemo()
	floatDemo()
}
