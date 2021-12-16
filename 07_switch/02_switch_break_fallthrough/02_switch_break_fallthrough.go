package main

import "fmt"

/*
	break：可以使用在switch和for循环中
		强制结束case语句，从而结束switch分支
	fallthrough：用于穿透switch
		当switch中某个case匹配成功后，就执行该case语句
		如果遇到fallthrough，则后面紧跟的case无需匹配，会直接穿透执行
		fallthrough只能位于在case的最后一行，否则报错
*/

func main() {
	n := 2
	switch n {
	case 1:
		fmt.Println("熊大")
		fmt.Println("熊大")
		fmt.Println("熊大")
	case 2:
		fmt.Println("熊二")
		fmt.Println("熊二")
		break // 强制结束这个case，一个case结束意味着switch结束了
		fmt.Println("熊二")
	case 3:
		fmt.Println("光头强")
		fmt.Println("光头强")
		fmt.Println("光头强")
	}

	fmt.Println("----------")

	m := 2
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	}

}
