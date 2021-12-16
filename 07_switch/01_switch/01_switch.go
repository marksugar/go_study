package main

import "fmt"

func main() {
	/*
		1、switch可以作用在其他类型上，但是case后的数值必须和switch作用的变量类型一致
		2、case是无序的，default通常放在最后
		3、case后的数值是唯一的，不能重复
		4、default非必须，是可选的
	*/

	num := 3
	switch num {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	default:
		fmt.Println("数据有误...")
	}

	fmt.Println("----------")

	// 模拟计算器
	num1 := 0
	num2 := 0
	oper := ""

	fmt.Println("请输入一个整数:")
	fmt.Scanln(&num1)
	fmt.Println("请输入一个整数:")
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
