/*
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2021-12-16 22:32:28
 * @LastEditTime: 2022-01-15 14:24:39
 * @version: v1.0
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test() {
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

func randDemo() {
	// time.now().Unix(): 返回一个从1970年1月1日0时0分0秒到现在得秒数
	// fmt.Println(time.Now().Unix())

	// 生成随机数还要设置一个种子
	// rand.Seed(time.Now().Unix())

	// unix 秒可能随机的还不够，使用纳秒来生成随机
	// rand.Seed(time.Now().UnixNano())

	// 生成100以内的随机数
	// n := rand.Intn(100) + 1 // [0,100)
	// fmt.Println(n)

	// 随机生成一个1-100的随机数，直到生成了99这个数，则退出，看看一共用了几次得到99这个数
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=",n)
		count++

		if n == 99 {
			break
		}
	}
	fmt.Printf("生成 99 这个数一共用了 %v 次", count)
}

func main() {
	// test()
	randDemo()
}
