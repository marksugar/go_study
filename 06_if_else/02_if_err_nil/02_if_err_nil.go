package main

import (
	"fmt"
	"strconv"
)

// 参考：https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/05.2.md

func main() {
	/*
		Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败
		当不使用 true 或 false 的时候，也可以使用一个 error 类型的变量来代替作为第二个返回值：成功执行的话，error 的值为 nil，
		否则就会包含相应的错误信息。Go 语言中的错误类型为 error: var err error。
		这样一来，就很明显需要用一个 if 语句来测试执行结果；由于其符号的原因，这样的形式又称之为 comma,ok 模式（pattern）
	*/

	/*
		函数 strconv.Atoi 的作用是将一个字符串转换为一个整数。
		之前我们忽略了相关的错误检查:anInt, _ = strconv.Atoi(origStr)
		如果 origStr 不能被转换为整数，anInt 的值会变成 0 而 _ 无视了错误，程序会继续运行。
		程序应该在最接近的位置检查所有相关的错误，至少需要暗示用户有错误发生并对函数进行返回，甚至中断程序。
		改进后程序如下：
	*/
	var orig string = "ABC"
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig [%s] is not an integer---exiting with error\n", orig)
	}

	fmt.Printf("The integer is: %d\n", an)

	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The newS strings is :%s\n", newS)

	/*
			这是测试 err 变量是否包含一个真正的错误（if err != nil）的习惯用法。如果确实存在错误，则会打印相应的错误信息然后通过 return 提前结束函数的执行。
			我们还可以使用携带返回值的 return 形式，例如 return err。这样一来，函数的调用者就可以检查函数执行过程中是否存在错误了。

		习惯用法：
		value, err := pack1.Function1(param1)
		if err != nil {
			fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
			return err
		}
		// 未发生错误，继续执行：
		由于本例的函数调用者属于 main 函数，所以程序会直接停止运行。
	*/

	//如果我们想要在错误发生的同时终止程序的运行，我们可以使用 os 包的 Exit 函数：
	//此处的退出代码 1 可以使用外部脚本获取到
	// 习惯用法：
	//if err != nil {
	//	fmt.Printf("Program stopping with error %v", err)
	//	os.Exit(1)
	//}

}
