/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-16 18:21:33
 * @LastEditTime: 2022-01-16 18:44:41
 * @version: v1.0
 */

/*
包名通常为和文件所在的目录同一个名字
但是Go中没有说包名必须要和目录的名字一模一样，但是为了方便管理，一般来说都是保持一致
尽可能使用约定俗成的好习惯
*/

package utils

import "fmt"

//为了让其他的包的文件使用Cal函数，需要将函数名首字母大写
// 这样类似于其他语言的public函数，理解为该函数是可以被公开给其他包的文件调用的函数
func Cal(n1 float64, n2 float64, operator byte) float64 {

	var res float64

	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符错误...")
	}
	return res
}
