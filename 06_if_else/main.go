package main

import (
	"fmt"
)

func main() {
	/*
	// 单 if 分支
	if condition {
		// do something	
	}
	
	// if ... else 分支
	if condition {
		// do something	
	} else {
		// do something	
	}
	
	// if...else if 分支
	if condition1 {
		// do something	
	} else if condition2 {
		// do something else	
	} else {
		// catch-all or default
	}
	
	*/

	salary := 10000
	if salary > 8000 {
		fmt.Println("工资大于8000")
	}

	height := 175
	if height < 175 {
		fmt.Println("身高低于1米75")
	} else if height > 175 {
		fmt.Println("身高大于1米75")
	} else {
		fmt.Println("身高等于1米75")
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

	// if其他写法.if 可以包含一个初始化语句（如：给一个变量赋值）,此方法涉及到作用域，以下变量num1只在这个if中有效
	if num1 :=4;num1 >0 {
		fmt.Printf("num1是%d,是正数\n",num1)
	}

	num2 := 10
	if num2 >0 {
		fmt.Printf("num2是%d,是正数\n", num2)
	}
	
	
	
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}
	
	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}
	
	// 通过在初始化语句中获取函数 process() 的返回值，并在条件语句中作为判定条件来决定是否执行 if 结构中的代码
	if value := process(data); value > max {
		...
	}
	
	
	/*
	这里不需要使用 if bool1 == true 来判断，因为 bool1 本身已经是一个布尔类型的值。
	可以使用取反 ! 来判断值的相反结果，如：if !bool1 或者 if !(condition)。后者的括号大多数情况下是必须的，如这种情况：if !(var1 == var2)。
	*/
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}
	
	/*
	当 if 结构内有 break、continue、goto 或者 return 语句时，Go 代码的常见写法是省略 else 部分
	无论满足哪个条件都会返回 x 或者 y
	*/
	if condition {
		return x
	}
	return y
	
	// 判断一个字符串是否为空
	if str == "" {
		...
	}
	if len(str) == 0 {
		...
	}
	
	// 通过常量 runtime.GOOS 来判断运行 Go 程序的操作系统类型.这段代码一般被放在 init() 函数中执行。
	 if runtime.GOOS == "windows"	 {
		.	..
	 } else { // Unix-like
		.	..
	 }
	
}


// 根据操作系统来决定输入结束的提示
 var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."

 func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
		fmt.Println("Is windows")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
		fmt.Println("Not windows")
	}
 }

// 函数 Abs() 用于返回一个整型数字的绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
theAbs := Abs(5)
fmt.Println(theAbs)

// isGreater 用于比较两个整型数字的大小
func isGreater(x, y int) bool {
	if x > y {
		return true	
	}
	return false
}
