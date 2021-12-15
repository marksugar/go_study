package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
格式化打印占位符：
%v,原样输出
%T,打印类型
%t,bool类型
%s,字符串类型
%f,浮点型
%d,10进制的整数
%b,2进制的整数
%o,8进制
%x,16进制，0-9,a-f
%X,16进制，0-9,A-F
%c,打印字符
%p,打印内存地址
 */

func main() {

	//a := 100
	//b := 4.25

	// 读取键盘输入
	//var x int
	//var y float64
	//fmt.Println("请输入一个整数和一个浮点数: ")
	//fmt.Scanln(&x, &y) // 读取键盘的输入，通过操作内存地址赋值给x和y。其运行时是阻塞式的
	//
	//fmt.Println("-----------")
	//
	//fmt.Printf("a的数值为：%d, b的数值为：%f\n", a, b)
	//
	//fmt.Println("==========")
	//fmt.Println("请输入一个整数和一个浮点数: ")
	//fmt.Scanf("%d,%f",&x,&y)
	//
	//fmt.Printf("x: %d, y:%f\n", x,y)

	// bufio包读取键盘输入
	fmt.Println("请输入一个字符串: ")
	reader := bufio.NewReader(os.Stdin)
	//s1, err := reader.ReadString('\n')  // 读到的标志就是到\n这里
	//if err != nil {
	//	return
	//}
	s1, _ := reader.ReadString('\n')  // 读到的标志就是到\n这里
	fmt.Println("读取到的键盘输入的数据为: ",s1)
}
