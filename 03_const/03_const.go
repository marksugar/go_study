package main

import "fmt"

func main() {
	fmt.Println(100)
	fmt.Println("allenjol")

	const URL string = "https://www.ayunw.cn"
	fmt.Println("URL: ", URL)

	const PI = 3.14
	fmt.Println("PI: ", PI)

	const c1, c2, c3 = 100, 3.14, "jol"
	fmt.Println("c1: ", c1, "c2: ", c2, "c3: ", c3)

	const (
		GENDER = "female"
		SALARY = 28000
		HOBBY  = "reading"
	)

	fmt.Println("GENDER: ", GENDER, "SALARY:", SALARY, "HOBBY:", HOBBY)

	const (
		a int = 100
		b
		c string = "jol"
		d
	)
	fmt.Printf("%T, %d\n", a, a)
	fmt.Printf("%T, %d\n", b, b)
	fmt.Printf("%T, %s\n", c, c)
	fmt.Printf("%T, %s\n", d, d)

	type scheme string
	const (
		H  scheme = "HTTP"
		HS scheme = "HTTPS"
	)
	fmt.Println("H: ",H, "HS:",HS)

	
	
	// 枚举类型 iota
	/*
	  iota是一个特殊的常量，可以被编译器自动修改的常量
	  每当定义一个const，iota第一次出现的时候为0
	  每当定义一个常量，iota就会自动增加1
	  直到下一个iota出现，值又变成0
	  
	  iota 在const 关键字出现的时候被重置为0，const 中的变量每声明一行常量声明将使iota 计数一次
	*/
	
	const (
		n1 = iota // n1=0
		n2        // n2=1
		n3        // n3=2
	)
	fmt.Println("n1: ",n1, "n2: ",n2, "n3: ",n3)

	const (
	n11 = iota // n1=0
	n12        // n2=1
	n13        // n3=2
	)
	fmt.Println("n11: ",n11, "n12: ",n12, "n13: ",n13)
	
	// 匿名变量跳过某些值
	const (
		m1 = iota //0
		m2        //1
		_
		m4 //3
	)
	fmt.Println("m1: ",m1, "m2: ",m2, "m4: ",m4)

	// iota声明中间插队
	const (
		j1 = iota //0
		j2 = 100  //100
		j3 = iota //2
		j4        //3
	)
	const j5 = iota //0

	fmt.Println("j1: ",j1, "j2: ",j2, "j3: ",j3, "j4: ",j4, "j5: ",j5)

	// 定义数量级
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println("KB: ",KB, "MB:",MB, "GB:",GB, "TB: ",TB, "PB: ",PB)

	// 多个iota定义在一行
	const (
		q, w = iota + 1, iota + 2 //1,2
		e, r                      //2,3
		t, y                      //3,4
	)
	fmt.Println("q: ",q, "w: ",w, "e:",e, "r: ",r, "t:",t, "y: ",y)

	// 示例1
	/*const (
	  // v1 = 1
	  // v2 =2
	  // v3 =3
	  // v4 =4
	  // v5 =5
	  )
	  fmt.Println(v1, v2, v3, v4, v5)
	*/

	// 示例2
	/*const (
	     v1 = iota  // 默认从 0 开始
	     v2
	     v3
	  )
	  fmt.Println(v1, v2, v3)
	*/

	// 示例3
	/*const (
	     v1 = iota + 1 // 从1开始
	     v2
	     v3
	  )
	  fmt.Println(v1, v2, v3)
	*/

	// 示例4
	/*const (
	     v1 = iota + 1
	     _  // 下划线这里可以作为一个常量存在，会获得iota+1以后的值，这里值为2
	     v2
	     v3
	  )
	  fmt.Println(v1, v2, v3)
	*/
}
