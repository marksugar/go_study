/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2021-12-17 21:19:35
 * @LastEditTime: 2022-01-12 21:46:40
 * @version: v1.0
 */

/*
	注意点：
		case后是一个表达式。即 常量、变量、一个有返回值的函数都可以
		case后的各个表达式的值的数据类型必须和switch的表达式数据类型一致
		case后面可以带多个表达式，使用逗号间隔，比如:case 表达式1,表达式2
		case后面的表达式如果是常量值(字面量)，则要求不能重复
		case后面不需要带break，程序匹配到一个case后就会执行对应的代码块，然后退出switch。如果一个都匹配不到则会执行default
		default不是必须的
		switch 不带表达式，类似if...else形式
*/

/*
	1、switch可以作用在其他类型上，但是case后的数值必须和switch作用的变量类型一致
	2、case是无序的，default通常放在最后
	3、case后的数值是唯一的，不能重复
	4、default非必须，是可选的
*/

package main

import "fmt"

func practiceDemo1() {
	// num := 3
	var num int
	fmt.Println("请输入一个月份:")
	fmt.Scanln(&num)

	switch num {
	case 1, 2, 3:
		fmt.Printf("输入的月份是%v月,是第一季度\n", num)
	case 4, 5, 6:
		fmt.Printf("输入的月份是%v月,是第二季度\n", num)
	case 7, 8, 9:
		fmt.Printf("输入的月份是%v月,是第三季度\n", num)
	case 10, 11, 12:
		fmt.Printf("输入的月份是%v月,是第四季度\n", num)
	default:
		fmt.Println("数据有误...")
	}

	fmt.Println("----------")

	// 模拟计算器
	num1 := 0
	num2 := 0
	oper := ""

	fmt.Println("请输入第一个整数:")
	fmt.Scanln(&num1)
	fmt.Println("请输入第二个整数:")
	fmt.Scanln(&num2)
	fmt.Println("请输入一个操作: +,-,*,/")
	fmt.Scanln(&oper)

	switch oper {
	case "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	}

	fmt.Println("----------")

	// 省略switch后的变量，相当于默认就是true
	switch {
	case true:
		fmt.Println("true...")
	case false:
		fmt.Println("flase...")
	}

	fmt.Println("----------")

	/* 成绩练习
	0-59 不及格，60-69 及格，70-79 中等，80-89 良好，90-100 优秀
	*/
	score := 88
	switch {
	case score >= 0 && score < 60:
		fmt.Println("不及格")
	case score >= 60 && score <= 70:
		fmt.Println("及格")
	case score >= 70 && score <= 80:
		fmt.Println("中等")
	case score >= 80 && score <= 90:
		fmt.Println("良好")
	case score >= 90 && score <= 100:
		fmt.Println("优秀")
	default:
		fmt.Println("成绩有误")
	}

	fmt.Println("----------")

	// case 后可以有多个表达式，用逗号隔开
	letter := "O"
	switch letter {
	case "A", "E", "I", "O", "U":
		fmt.Println(letter, "是元音")
	case "M", "N":
		fmt.Println(letter, "M或者N")
	default:
		fmt.Println("其他...")
	}

	fmt.Println("----------")

	// 判断一个月的天数
	month := 5
	day := 0
	year := 2021
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		fmt.Println("月份有误")
	}
	fmt.Printf("%d 年 %d 月的天数是: %d\n", year, month, day)

	fmt.Println("----------")

	// switch后可以多一条初始化语句，变量作用域就只在switch中
	// 语法：switch 初始化语句;变量
	switch language := "golang"; language {
	case "golang":
		fmt.Println("Go语言")
	case "java":
		fmt.Println("Java语言")
	case "python":
		fmt.Println("Python语言")
	}
}

func test(char byte) byte {
	return char + 1
}

func practiceDemo2() {
	var str byte
	fmt.Println("请输入一个字符[a,b,c,d]")
	fmt.Scanf("%c", &str)

	switch test(str) + 1 {
	case 'a':
		fmt.Println("哞哞")
	case 'b':
		fmt.Println("布谷")
	case 'c':
		fmt.Println("汪汪")
	case 'd':
		fmt.Println("滴滴滴")
	default:
		fmt.Println("数据有误...")
	}
}
func practiceDemo3() {
	var n1 int32 = 5
	var n2 int32 = 10
	var n3 int32 = 5

	switch n1 {
	case n2, 10, 5:
		fmt.Println("ok1")
	case n3: // n3是个变量而不是常量值5，因此编译器是允许这样的。如果直接写5则编译器会报错，印上上一个case已经有常量5了
		fmt.Println("ok2...")
	default:
		fmt.Println("没正确匹配")
	}

}

func practiceDemo4() {
	// 类似if else 的方式来用
	var age int = 20
	switch {
	case age == 10:
		fmt.Println("age = 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没匹配到")
	}
	fmt.Println("----------")

	// case 中对score 进行范围判断
	var score float64 = 95
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score >= 70 && score <= 90:
		fmt.Println("优良")
	case score >= 60 && score <= 70:
		fmt.Println("及格")
	default:
		fmt.Println("没匹配到")
	}

	// switch 后可以直接声明一个变量，用分号结束。不推荐这样玩
	switch grade := 90; {
	case grade > 90:
		fmt.Println("优秀...")
	case grade >= 70 && grade <= 90:
		fmt.Println("优良...")
	case grade >= 60 && grade <= 70:
		fmt.Println("及格...")
	default:
		fmt.Println("没匹配到...")
	}
}

func main() {
	//practiceDemo1()
	//practiceDemo2()
	//practiceDemo3()
	practiceDemo4()
}

