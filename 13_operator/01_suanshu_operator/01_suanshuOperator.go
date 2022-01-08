package main

import "fmt"

func main() {
	// 假如还有97天放假，问：还有几个星期零几天
	var holiday int = 97
	var weeks int = holiday / 7
	var days int = holiday % 7

	fmt.Printf("还有%d个星期零%d天\n", weeks, days)

	// 定义一个变量保存华氏度，华氏度转换摄氏度的公式为
	// 5/9*(华氏度-100)，求出华氏温度对应的摄氏温度
	var huashi float64 = 134.2
	var sheshi float64 = 5.0 / 9 * (huashi - 100)

	fmt.Printf("华氏度=%v，对应的摄氏度=%v\n", huashi, sheshi)

}
