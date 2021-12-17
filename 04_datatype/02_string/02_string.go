package main

import (
	"fmt"
	"strings"
)

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
			解释字符串：
				\n 换行
				\t 制表符，相当于tab键
				\r 回车符
				\u或者\U Unicode字符
				\\ 反斜杠自身
			还有一个特例是输出百分号，需要用两个百分号表示：%%

		非解释字符串：使用反引号括起来，支持换行，如：
			`This is a raw string \n` 这句话中的`\n\` 会被原样输出
	*/

	fmt.Println("HelloWorld")
	fmt.Println("\"HelloWorld\"")
	fmt.Println("Hello\nWor\tld")
	fmt.Println(`Hel"loWor"ld`)
	fmt.Println("Hello`Wor`ld")

	// 用两个反引号 `` 包裹的内容会原样输出
	sss := `春眠不觉晓
		一觉睡到老，处处都有鸟
				夜来风雨声，花落知多少`
	fmt.Println(sss)

	fmt.Println("----------")

	// 字符串拼接
	str1 := "allen"
	str2 := "jol"
	str := str1 + str2
	fmt.Println("str:", str)

	fmt.Println("----------")

	// 由于编译器行尾自动补全分号的缘故，加号 + 必须放在第一行
	str5 := "Beginning of the string " +
		"second part of the string"
	fmt.Println("str2：", str5)

	// 简写形式 += 拼接字符串
	str6 := "hel" + "lo,"
	str6 += "world"
	fmt.Println("str6:", str6)

	/*
		注意：
		用 + 号拼接字符并不是最搞笑的方法，更好的办法是用函数strings.Join()
		还有更好的办法是使用字节缓冲bytes.Buffer 拼接更给力
	*/

	/*
		strings.Join() 用法：
			func Join(s []string, sep string) string
			其中s是可以用来连接元素的字符串，sep是放在字符串元素之间的分隔符
	*/
	str7 := []string{"jenkins", "gitlab", "k8s", "envoy"}
	fmt.Println(strings.Join(str7, "-"))

	str8 := []string{"ayunw", "cnsre", "google", "iyunw"}
	fmt.Println(strings.Join(str8, " "))
}

