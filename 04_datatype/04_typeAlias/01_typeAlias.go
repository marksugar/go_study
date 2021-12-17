package main

import "fmt"

// 声明一个类型为int的别名TZ
type TZ int

func main() {
	var a, b TZ = 4, 5
	c := a + b
	fmt.Printf("The value of c is: %d\n", c)

	// 定义一个类型为string的别名Rope，并声明一个该类型的变量
	type Rope string
	var name Rope = "allen_jol"
	fmt.Printf("name is: %s\n",name)
}
