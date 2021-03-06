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

	fmt.Println("f5=", f5, "\t", "f6=", f6)

	num1 := 5.12
	num2 := .123
	fmt.Println("num1=", num1, "num2=", num2)

	num3 := 5.1234e2 // 5.1235 * 10的2次方
	num4 := 5.1234e2
	num5 := 5.1234e-2 // 5.1235 / 10的2次方
	fmt.Println("num3=", num3, "num4=", num4, "num5=", num5)
}

func stringDemo() {
	/*
		go的字符串是由字节组成的
		1、如果保存的字符在ASCII码表中，如：[0-9,a-z,A-Z]，则可以用byte
		2、如果保存的字符对应的码值大于255，则需要考虑int类型保存
		3、如果需要输出字符，则需要用格式化输出，使用%c来表示
		4、go语言的字符使用UTF-8编码
	*/
	fmt.Println("----------")

	c1 := 'a'
	c2 := '0'

	fmt.Println("c1=", c1)
	fmt.Printf("c1的类型: %T\n", c1)

	fmt.Println("c2=", c2)
	fmt.Printf("c2的类型：%T\n", c2)
	// %c初始的是对应的字符
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	fmt.Println("----------")

	var d1 byte = 'a'
	var d2 byte = '0'
	fmt.Println("d1=", d1)
	fmt.Printf("d1的类型: %T\n", d1)
	fmt.Println("d2=", d2)
	fmt.Printf("d2的类型：%T\n", d2)

	//var d3 byte = '南'  // byte范围是0~255，因此这里overflow 溢出。
	var d3 int = '南' // int类型的范围
	fmt.Printf("d3=%c d3的码值=%d\n", d3, d3)

	var d4 = "东"
	fmt.Printf("d4=%s\n", d4)

	// 字符类型可以进行运算，相当于整数
	var e1 = 10 + 'a' // 小a的ASCII码为 97。运算时会根据ASCII值计算
	fmt.Println("e1=", e1)

	// go中字符串一旦赋值就不能更改
	//str1 := "hello"
	//str1[0] = "m"  // go中字符串不可变

	// 双引号会识别转义符，反引号可以原生输出字符
	str2 := "abc\nefg"
	fmt.Println(str2)

	str3 := `
	package main
	import (
		"fmt"
	)
	func main() {
		var s1 string = "tree"
		fmt.Println("s1=", s1)
	}
	`
	fmt.Println(str3)

	// 字符串拼接
	var str4 = "hello" + " golang"
	str4 += " let's go"
	fmt.Println(str4, "\n")

	// 字符串拼接的操作很长，必须将加号留在上一行
	var str5 = "hello" + " golang" +
		" let's go" + " to your" +
		" favorite provisioner"

	fmt.Println(str5)
}

func defaultDataType(){
	fmt.Println("----- default data type -----")
	var a int			// 0
	var b float32		// 0
	var c float64		// 0
	var isMarried bool	// false
	var name string		// ""
	// %v 表示按照变量的值输出
	fmt.Printf("a=%d,b=%f,c=%f,isMaried=%v,name=%v", a,b,c,isMarried,name)


}

func main() {
	intDemo()
	floatDemo()
	stringDemo()
	defaultDataType()
}

