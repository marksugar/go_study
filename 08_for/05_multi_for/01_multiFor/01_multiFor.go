/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-13 22:39:54
 * @LastEditTime: 2022-01-15 10:20:29
 * @version: v1.0
 */

package main

import "fmt"

/*
题目：
	学生的成绩从键盘输入
	1、统计三个班的成绩，一个班级5名学生，求出各个班的平均分和所有班级的平均分
	2、统计三个班级的学生成绩及格的人数
*/

/*
思路1
	1、统计1个班的成绩，一个班级5个学生，求出该班的平均分，学生成绩从键盘输入
	2、i表示学生个数
	3、声明一个sum来统计班级的总分

思路2
	1、在思路1的基础上统计 3 个班的成绩，一个班级5个学生，求出每个班的平均分，学生成绩从键盘输入
	2、 j表示第几个班级

思路3
	1、定义两个变量，表示班级的个数和班级的人数

思路4
	1、定义变量paasCount 保存及格的人数个数
*/

func studentScores() {
	var (
		classNum   int = 3
		studentNum int = 5
	)

	var totalSum float64 = 0.0 // 3个班级的总成绩的变量
	var paasCount int = 0
	for class := 1; class <= classNum; class++ { // 循环班级
		var sum float64 = 0.0                                   // 一个班的学生总分
		for students := 1; students <= studentNum; students++ { // 循环输入5个学生的分数
			var score float64 // 单个学生的分数
			fmt.Printf("请输入第%d个班的第%d个学生的成绩: ", class, students)
			fmt.Scanln(&score)
			sum += score // 1 个班级的学生总分数
			// 判断并且统计分数及格的分数
			if score >= 60 {
				paasCount++
			}
		}

		fmt.Println()
		fmt.Printf("第%d个班级的平均分是:%v\n", class, sum/float64(studentNum))
		// 把各个班的总成绩放到 totalSum 中
		totalSum += sum
	}

	fmt.Println()
	fmt.Printf("各个班级的总分:%v,所有班级的平均分是:%v\n", totalSum, totalSum/float64(studentNum*classNum))

	fmt.Println()
	fmt.Printf("%d个班级的成绩及格的人数:%v\n", classNum, paasCount)
}

func main() {
	studentScores()
}
