package main

import "fmt"

/*
	&&  逻辑与，也叫短路与。如果第一个条件为false，则第二个条件不会判断，最终结果为false
	||	逻辑或，也叫短路或。如果第一个条件为true，则第二个条件不会判断，最终结果为true
	！	逻辑非
*/

func logic_operator() {
	// 逻辑与 &&
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("yes, && ok1")
	}

	if age > 30 && age <= 40 {
		fmt.Println("yes, && ok2")
	}

	// 逻辑或 ||
	if age > 30 || age < 50 {
		fmt.Println("yes, && ok3")
	}

	if age > 30 || age <= 40 {
		fmt.Println("yes, && ok4")
	}

	// 逻辑非 !，可以理解为将结果取反
	if age > 30 {
		fmt.Println("yes, && ok5")
	}

	if !(age > 30) { // age > 30 为true,结果取反后为false则不会输出
		fmt.Println("yes, && ok6")
	}

}

func test() bool {
	fmt.Println("test...")
	return true
}

func main() {
	logic_operator()

	var i int = 10
	if i > 9 && test() {
		fmt.Println("i > 9,ok...")
	}

	// i < 9为false，因此后面test()函数不会被判断，就没有任何输出
	if i < 9 && test() {
		fmt.Println("i < 9,ok...")
	}

	// i < 9 为false，则或判断test()函数，因此整个结果为true，会输出test()函数中的值和if中的打印结果
	if i < 9 || test() {
		fmt.Println("i < 9,or ok...")
	}

	// 因为 i > 9为true，因此不会再去判断test()函数了，只会输出if判断中的打印
	if i > 9 || test() {
		fmt.Println("i > 9,or ok...")
	}
}
