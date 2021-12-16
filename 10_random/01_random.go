package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		随机数操作都在math/rand包下
	*/
	num := rand.Int()
	fmt.Println(num)

	for i := 0; i < 10; i++ {
		num2 := rand.Intn(10) // 0~9之间
		fmt.Println(num2)
	}
	rand.Seed(10000)
	num3 := rand.Intn(10)
	fmt.Println("随机数num3:", num3)

	fmt.Println("----------")

	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%T\n", t1)

	// 时间戳：距离1970年1月1日0点0分0秒，之间的时间差值：秒，纳秒
	ts := t1.Unix()
	fmt.Println("timestamp:", ts)

	tsnano := t1.UnixNano()
	fmt.Println("timestampNano:", tsnano)

	fmt.Println("----------")

	/*
		用时间戳作为随机数的种子数
	*/

	//step1、用时间戳设置种子数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		// 调用生成随机数的方法
		fmt.Println(rand.Intn(100))
	}


	fmt.Println("----------")

	/* 获取指定范围内的随机数
	Intn  [0,n)

	[15,76]
		[0,76-15] 也就是 [0,61] + 15
	*/

	num4 := rand.Intn(61) + 15
	fmt.Println("num3的值：", num4)
}
