package main

import "fmt"

/*
计算机编码：
	计算机本质只能识别0和1
	A: 65, B:66, C:67 ...
	a: 97, b:98, c:99 ...
中国编码表：gbk，兼容ASCII码
Unicode编码：UTF-8, UTF-16，UTF-32...

Go语言统一使用UTF-8编码
*/

func main() {
	s1 := "王二"
	fmt.Printf("%T, %s\n", s1, s1)

	s2 := "张三"
	fmt.Printf("%T,%s\n", s2, s2)

	// 'A' 和 "A" 的区别：编码上的区别
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%s\n", v1, v1) // int32
	fmt.Printf("%T,%s\n", v2, v2)

	v3 := "中"
	fmt.Printf("%T, %d, %c, %q\n", v3, v3, v3, v3)

	// 字符转义,使用 \
	/*
		转义成普通字符：\'  ,  \"
		特殊：
			\n 换行
			\t 制表符
	 */
	fmt.Println("HelloWorld")
	fmt.Println("\"HelloWorld\"")
	fmt.Println("Hello\nWor\tld")
	fmt.Println(`Hel"loWor"ld`)
	fmt.Println("Hello`Wor`ld")

}
