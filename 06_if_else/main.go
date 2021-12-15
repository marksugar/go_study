package main

import (
	"fmt"
)

func main() {

	salary := 10000
	if salary > 8000 {
		fmt.Println("工资大于8000")
	}

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

	//sex := "泰国人妖"
	//if sex == "男" {
	//	fmt.Println("前面左拐男厕所...")
	//} else if sex == "女" {
	//	fmt.Println("前面右拐女厕所...")
	//} else {
	//	fmt.Println("性别非男非女，怕是人妖...")
	//}

	// if 语句嵌套
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

	fmt.Println("----------")

	// if其他写法.涉及到作用域
	if num1 :=4;num1 >0 {
		fmt.Printf("num1是%d,是正数\n",num1)
	}

	num2 := 10
	if num2 >0 {
		fmt.Printf("num2是%d,是正数\n", num2)
	}
}
