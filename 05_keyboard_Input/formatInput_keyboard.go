package main

import "fmt"

/*
document: https://studygolang.com/pkgdoc

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

Scanf或者Scanln 从标准输入扫描文本
*/

var (
	name   string
	age    int
	salary float64
	isPass bool
)

func demo1() {
	//a := 100
	//b := 4.25

	// 读取键盘输入
	//var x int
	//var y float64
	//fmt.Println("请输入一个整数和一个浮点数: ")
	//fmt.Scanln(&x, &y) // 读取键盘的输入，通过操作内存地址赋值给x和y。其运行时是阻塞式的

	//fmt.Println("-----------")

	//fmt.Printf("a的数值为：%d, b的数值为：%f\n", a, b)

	//fmt.Println("==========")
	//fmt.Println("请输入一个整数和一个浮点数: ")
	//fmt.Scanf("%d,%f",&x,&y)

	//fmt.Printf("x: %d, y:%f\n", x,y)

	// bufio包读取键盘输入
	//fmt.Println("请输入一个字符串: ")
	//reader := bufio.NewReader(os.Stdin)
	////s1, err := reader.ReadString('\n')  // 读到的标志就是到\n这里
	////if err != nil {
	////	return
	////}
	//s1, _ := reader.ReadString('\n') // 读到的标志就是到\n这里
	//fmt.Println("读取到的键盘输入的数据为: ", s1)
}

func scanfDemo() {
	// 从控制台接收键盘输入的用户信息：姓名、年龄、工资、是否通过公司考核
	// 1. fmt.Scanln()
	//var (
	//	name   string
	//	age    int
	//	salary float64
	//	isPass bool
	//)
	fmt.Println("请输入姓名: ")
	// 当程序执行到 fmt.Scanln(&name) 的时候，程序会停止在这里等待用户输入，并且回车
	fmt.Scanln(&name)

	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)

	fmt.Println("请输入工资: ")
	fmt.Scanln(&salary)

	fmt.Println("请输入是否通过考试[true|false]: ")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是: \033[1;31m%v\033[0m \n年龄是: \033[1;31m%v\033[0m \n工资是: \033[1;31m%v\033[0m \n是否通过公司考核: \033[1;31m%v\033[0m\n", name, age, salary, isPass)

	fmt.Println("--------------------")
}

func scanlnDemo(){
	// 2. fmt.Scanf(),可以按照指定格式输入
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过公司考核， 输入时请使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("名字是: \033[1;32m%v\033[0m \n年龄是: \033[1;32m%v\033[0m \n工资是: \033[1;32m%v\033[0m \n是否通过公司考核: \033[1;32m%v\033[0m\n", name, age, salary, isPass)
}

func main() {
	demo1()
	scanfDemo()
	scanlnDemo()
}
