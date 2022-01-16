/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-16 18:21:42
 * @LastEditTime: 2022-01-16 18:54:55
 * @version: v1.0
 */

package main

import (
	"fmt"
	// 一定要从src的下一层目录开始导入，一直写到要引入的包的最后一层的目录
	// import引入包时，路径从$GOPATH的src下开始，不用写src，因为编译器会自动从src下开始引入
	"github.com/allenjol/16_package/utils"
	// 给包起别名为util，那么所有调用utils这个包下的函数的时候，调用的时候函数都要写util而不是utils
	// 也就是说在调用的时候要用别名来访问该包的函数和变量
	// util "github.com/allenjol/16_package/utils"
)

func main() {
	var (
		n1       float64 = 1.2
		n2       float64 = 2.3
		operator byte    = '+'
	)

	result := utils.Cal(n1, n2, operator)
	fmt.Println("赋初始值的result=", result)

	// 给n1, n2, operator重新赋值后打印结果
	n1 = 4.5
	n2 = 6.7
	operator = '*'

	result = utils.Cal(n1, n2, operator)
	fmt.Printf("重新赋值后的result=%.3f", result)
}
