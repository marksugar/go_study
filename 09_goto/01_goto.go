/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2021-12-16 22:22:38
 * @LastEditTime: 2022-01-16 15:02:52
 * @version: v1.0
 */
package main

import "fmt"

/*
goto语句可以无条件转移到程序中指定的位置
通常与条件语句配合使用。可以用来实现条件转移，跳出循环等功能

语法：
	goto label
	...		// 上面使用了togo以后，这部分的代码就不会执行了
	label: statement	// 直接跳转到此处

注意：Go中一般不主张使用goto，以免造成程序流程的混乱，导致理解和调试程序产生困难
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
			fmt.Println("j=", j)
			goto onExit
		}
	}
	return

onExit:
	fmt.Println("err here,exit...")

}

