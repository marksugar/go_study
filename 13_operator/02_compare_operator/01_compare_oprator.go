package main

import "fmt"

/*
关系运算符的结果是bool类型，true或者false
关系运算符组成的表达式叫关系表达式：a > b
比较运算符"==" 不要写成"="。两个等号代表判断是否相等，一个等号代表赋值
关系运算符一般用于if结构或者是循环结构
*/

func compare_operator() {
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2) // false
	fmt.Println(n1 != n2) // true
	fmt.Println(n1 > n2)  // true
	fmt.Println(n1 >= n2) // true
	fmt.Println(n1 < n2)  // false
	fmt.Println(n1 <= n2) // false

	flag := n1 > n2
	fmt.Println("flag=", flag) // true
}

func main() {
	compare_operator()
}
