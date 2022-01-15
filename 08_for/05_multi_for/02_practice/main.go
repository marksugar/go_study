/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-15 10:22:09
 * @LastEditTime: 2022-01-15 11:12:31
 * @version: v1.0
 */

package main

import "fmt"

/*
打印矩形
*/

func juxing() {
	// i表示矩形的层数
	for i := 1; i <= 5; i++ {
		// j 表示每层打印几个 * 号
		for j := 1; j <= 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
*	第一层1个 *
**	第二层2个 *
***	第三层3个 *
 */
func halfOfJinzita() {
	// i 控制金字塔的层数，j控制金字塔的每一层的 * 的个数
	for i := 1; i <= 3; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
打印金字塔
	1、定义一个变量表示金字塔的层数(行数)
	2、定义一个变量表示每层金字塔的 * 号的个数
	3、定义一个变量表示金字塔每一层的空格

	*		// i =1,j=1,space=3
   ***		// i=2,j=3,space=2
  *****		// i=3,j=5,space=1
 *******	// i=4,j=7,space=0

*/
func jinzita() {
	// 代表金字塔的层数
	var totalLevel int = 10

	for i := 1; i <= totalLevel; i++ {
		// 在打印 * 之前还要先打印空格
		for space := 1; space <= totalLevel-i; space++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
打印空心金字塔
	在打印金字塔的每一层的时候要考虑是打印 * 还是打印空格
	每层的第一个和最后一个打印 *，其他的就是打印空格
	最后一层(最底下这一层)是要全部打上 * 号的

	*
   * *
  *   *
 *     *
*********

*/
func kongxinJinzita() {
	// 代表金字塔的层数
	var totalLevel int = 5

	// i 表示金字塔的层数
	for i := 1; i <= totalLevel; i++ {
		// 在打印 * 之前还要先打印空格
		for space := 1; space <= totalLevel-i; space++ {
			fmt.Print(" ")
		}

		// j 表示每层打印多少个 *
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 九九乘法表
func chengfabiao() {
	var digitalNum int = 9

	for i := 1; i <= digitalNum; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		
		//表示换行
		fmt.Println()
	}
}

func main() {
	// juxing()
	// halfOfJinzita()
	// jinzita()
	// kongxinJinzita()
	chengfabiao()
}
