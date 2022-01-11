/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2021-12-15 22:52:02
 * @LastEditTime: 2022-01-11 21:56:01
 * @version: v1.0
 */

package main

import (
	"fmt"
	_ "math" // 暂时没用到math包，所以通过一个匿名变量来丢弃math包
)

func ifDemo() {
	// 单分支
	salary := 10000
	if salary > 8000 {
		fmt.Println("工资大于8000")
	}

	// 多分支
	score := 95
	if score < 60 {
		fmt.Println("成绩不及格")
	} else if score < 80 {
		fmt.Println("成绩良")
	} else if score < 90 {
		fmt.Println("成绩优")
	} else {
		fmt.Println("成绩非常优秀")
	}

	height := 175
	if height < 175 {
		fmt.Println("身高低于1米75")
	} else if height > 175 {
		fmt.Println("身高大于1米75")
	} else {
		fmt.Println("身高等于1米75")
	}

	sex := "泰国人妖"
	if sex == "男" {
		fmt.Println("前面左拐男厕所...")
	} else if sex == "女" {
		fmt.Println("前面右拐女厕所...")
	} else {
		fmt.Println("性别非男非女，怕是人妖...")
	}

	fmt.Println("----------")

	// if其他写法.涉及到作用域
	if num1 := 4; num1 > 0 {
		fmt.Printf("num1是%d,是正数\n", num1)
	}

	num2 := 10
	if num2 > 0 {
		fmt.Printf("num2是%d,是正数\n", num2)
	}
}

func ifDemoScanf() {
	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)

	// 双分支
	if age > 18 {
		fmt.Printf("年龄是%d岁,已经成年...\n", age)
	} else {
		fmt.Printf("年龄是%d岁,未成年...\n", age)
	}
}

func practiceDemo1() {
	var year int = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("%v年是闰年...\n", year)
	}
}

func practiceDemo2() {
	var b bool = true
	if b == false { // 如果是if b = false，则go中会编译出错。if的条件表达式不能是赋值语句
		fmt.Println("a")
	} else if b {
		fmt.Println("b")
	} else if !b {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}
}

// 从终端输入三个条件的值
func practiceDemo3() {
	var (
		height  float64
		money   float64
		hansome bool
	)

	fmt.Println("请输入你的身高:")
	fmt.Scanln(&height)
	fmt.Println("请输入你的存款:")
	fmt.Scanln(&money)
	fmt.Println("请输入帅气与否(true|false):")
	fmt.Scanln(&hansome)

	if height > 180 && money > 500000 && hansome == true {
		fmt.Println("高富帅，我想嫁给他!")
	} else if height > 180 || money > 500000 || hansome == true {
		fmt.Println("满足其中的一个条件，可以嫁给他!")
	} else {
		fmt.Println("一个条件都不满足，我继续做单身狗!")
	}
}

func practiceDemo4() {
	// if分支语句嵌套
	sex := "泰国人妖"
	if sex == "男" {
		fmt.Println("前面左拐男厕所...")
	} else {
		if sex == "女" {
			fmt.Println("前面右拐女厕所...")
		} else {
			fmt.Println("性别非男非女，怕是人妖...")
		}
	}
}

func practiceDemo5() {
	var (
		second float64
		gender string
	)

	fmt.Println("请输入秒数:")
	fmt.Scanln(&second)

	if second <= 8 {
		fmt.Println("小于等于8秒钟")
		fmt.Println()

		fmt.Println("请输入性别:")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入男子组决赛")
		} else {
			fmt.Println("进入女子组决赛")
		}
	} else {
		fmt.Println("超过了8秒钟，没有人能进入决赛")
	}
}

/*
淡旺季出票系统：
	旺季 4~10月份
		成年人(18-60): 60元
		儿童(<18): 成年人的半价门票
		老人(>60): 成年人门票价格的三分之一
	淡季：
		成年人 40，其他 20
*/

func piaoSystem() {
	var (
		month int
		age   int
		price float64 = 60.0
	)

	fmt.Println("请输入月份: ")
	fmt.Scanln(&month)
	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("淡季%v月,老人门票: %v\n", month, price/3)
		} else if age >= 18 {
			fmt.Printf("%v月,成年人门票: %v\n", month, price)
		} else {
			fmt.Printf("%v月,儿童门票: %v\n", month, price/2)
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Printf("淡季%v月,成年人门票: 40", month)
		} else {
			fmt.Printf("淡季%v月,其他门票: 20", month)
		}
	}

	//if month >=4 && month <= 10 {
	//	fmt.Println("请输入买票者年龄:")
	//	fmt.Scanln(&age)
	//	if age >18 && age <60 {
	//		fmt.Println("门票是60元")
	//	} else if age <18 {
	//		fmt.Println("门票是30元")
	//	} else if age > 60 {
	//		fmt.Println("门票是20元")
	//	} else {
	//		fmt.Println("输入年龄有误，请重新输入")
	//	}
	//} else if month<4 && month > 10 {
	//	fmt.Println("请输入买票者年龄:")
	//	fmt.Scanln(&age)
	//
	//	if age >=18 && age <= 60 {
	//		fmt.Println("门票为60元")
	//	} else if age <18 || age > 60 {
	//		fmt.Println("门票是20元")
	//	} else {
	//		fmt.Println("输入年龄有误")
	//	}
	//} else {
	//	fmt.Println("输入月份有误")
	//}
}

func main() {
	//ifDemo()
	//ifDemoScanf()
	//practiceDemo1()
	//practiceDemo2()
	//practiceDemo3()
	//practiceDemo4()
	//practiceDemo5()
	piaoSystem()
}

