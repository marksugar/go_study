package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）

		参考：
			https://www.cnblogs.com/zxqblogrecord/p/13303241.html
			https://www.jianshu.com/p/8ac396850a3d
	
		time包提供了一个数据类型time.Time作为值使用，以及显示和测量时间与日期的功能函数
		当前时间通过time.Now()获取，或者用t.Day()、t.Minute()等来获取时间的某一部分。
		自定义时间格式化字符：fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) 将会输出 18.12.2011

		Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。
		Location 类型映射某个时区的时间，UTC 表示通用协调世界时间
	*/

	var week time.Duration
	t := time.Now()
	fmt.Println("now time:", t)                                  // now time: 2021-12-18 13:52:56.744334 +0800 CST m=+0.000090079
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year()) // 18.12.2021

	t = time.Now().UTC()
	fmt.Println("utc time:", t) // utc time: 2021-12-18 05:52:56.744498 +0000 UTC
	fmt.Println(time.Now())     // 2021-12-18 13:52:56.7445 +0800 CST m=+0.000256334

	// 时间计算
	week = 60 * 60 * 24 * 7 * 1e9 // must be a nanosec
	weekFromNow := t.Add(time.Duration(week))
	fmt.Println(weekFromNow) // 2021-12-25 05:52:56.744498 +0000 UTC

	// 格式化时间
	fmt.Println(t.Format(time.RFC822)) // 18 Dec 21 05:52 UTC
	fmt.Println(t.Format(time.ANSIC))  // Sat Dec 18 05:52:56 2021

	// The time must be 2006-01-02 15:04:05，这是Go的诞生时间
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 18 Dec 2021 05:52
	s := t.Format("20060102")
	fmt.Println(t, "--->", s) // 2021-12-18 05:52:56.744498 +0000 UTC ---> 20211218

	// 暂停3秒钟
	fmt.Println("接下去先暂停3秒")
	time.Sleep(3 * time.Second)

	// 如果不是数字变量 则必须写上 time.Duration,如下
	//time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("暂停3秒结束")

	fmt.Println("----------")

	//获取年、月、日、时、分、秒
	getYear := time.Now().Year()        //获取年
	getMonth := time.Now().Format("01") //获取月
	getDay := time.Now().Day()          // 获取日

	// 或者

	getYear2 := time.Now().Format("2006") // 获取年
	getMonth2 := time.Now().Format("01")  // 获取月
	getDay2 := time.Now().Format("02")    // 获取日
	hour := time.Now().Format("15")
	min := time.Now().Format("04")
	second := time.Now().Format("05")

	fmt.Println("year:", getYear, "month:", getMonth, "day:", getDay)
	fmt.Println("第二章输出的时间方式如下...")
	fmt.Println("year:", getYear2, "month:", getMonth2, "day:", getDay2)
	fmt.Println("输出时分秒...")
	fmt.Println("hour:", hour, "min:", min, "second:", second)

	fmt.Println("----- 获取日期的各种格式 -----")

	//获取日期各种格式
	todayStr1 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("todaystr1:", todayStr1)

	todayStr2 := time.Now().Format("2006/01/02 15:04:05")
	fmt.Println("todaystr2:", todayStr2)

	todayStr3 := time.Now().Format("2006-01-02")
	fmt.Println("todaystr3:", todayStr3)

	todayStr4 := time.Now().Format("15:04:05")
	fmt.Println("todaystr4:", todayStr4)

	fmt.Println("----- 获取前一天的时间-----")
	// 获取前一天的日期
	currentTime := time.Now()
	theYesterdayTime := currentTime.AddDate(0, 0, -1) // 前一天就是-1，前5天就是-5
	fmt.Println("前一天的日期是:", theYesterdayTime)


	/*
	获取时间戳
	时间戳是自 1970 年 1 月 1 日（08:00:00GMT）至当前时间的总毫秒数。它也被称为 Unix 时 间戳（UnixTimestamp）。
	 */
	fmt.Println("----- 获取时间戳 -----")
	timeObj := time.Now()
	unixTime :=timeObj.Unix()	// 获取当前的时间戳(到秒)
	fmt.Println("当前到秒的时间戳：",unixTime)

	unixNatime :=timeObj.UnixNano()	// 获取当前的时间戳(到纳秒)
	fmt.Println("当前到纳秒的时间戳：",unixNatime)


	// 时间戳转换成字符串,假设用到秒的时间戳：1639808130
	fmt.Println("----- 时间戳转换成字符串 -----")
	unixTime2 := 1639808130
	timeObj2 :=time.Unix(int64(unixTime2), 0)
	fmt.Println("timeObj2:", timeObj2)

	var str = timeObj2.Format("2006-01-02 15:04:05")
	fmt.Println("从时间戳转换回来的时间:",str)

}
