package main

import "fmt"

/*
1、goto可以无条件的跳转到程序中指定的行
2、goto语句通常与条件语句配合使用，可以用来实现条件转移、跳出循环体等功能
3、在Go语言中一般不主张使用goto，能不用goto就不要用goto，容易引起程序混乱
*/
func main() {
	/*
		当a为15的时候,if判断成立，执行到goto，则会跳出if判断语句块后面的打印和a++
		然后又回到了for循环，因此不会打印15
	*/

	a := 10
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Printf("a的值为:%d\n", a)
		a++
	}

	// 统一错误处理,避免多处错误存在代码重复问题
	fmt.Println("----------")

	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			goto onExit
		}
	}
	return

onExit:
	fmt.Println("err here,exit...")

}
