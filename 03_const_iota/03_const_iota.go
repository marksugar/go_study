package main

import "fmt"

/*
	const来定义常量，变量定义后必须要使用，否则编译的时候报错
	而常量定义后可以不使用，编译时不会报错
	
	每遇到一次 const 关键字，iota 就重置为 0
*/

func main() {
	fmt.Println(100)
	fmt.Println("allenjol")

	const URL string = "https://www.ayunw.cn"
	fmt.Println("URL: ", URL)

	const PI = 3.14
	fmt.Println("PI:", PI)

	const c1, c2, c3 = 100, 3.14, "jol"
	fmt.Println("c1:", c1, "c2:", c2, "c3:", c3)
	
	// 使用 \ 来作为连接符
	const Ln2 = 0.693147180559945309417232121458\
			176568075500134360255254120680009
	const Log2E = 1/Ln2
	const Billion = 1e9
	const hardEight = (1 << 100) >> 97

	// 多个常量定义在一行，类型推导。不推荐这样使用，常量多了看的不清晰。
	const beef, two, third = "eat", 2, "veg"
	const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	const (
		Monday, Tuesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	
	// 批量定义常量，类型推导
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
	
	// iota 可以被用作枚举值
	const (
	n11 = iota  // n1=0
	n12 = iota  // n2=1
	n13 = iota  // n3=2
	)
	fmt.Println("n11: ",n11, "n12: ",n12, "n13: ",n13)
	
	// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
	const (
		n1 = iota // n1=0
		n2        // n2=1
		n3        // n3=2
	)
	fmt.Println("n1: ",n1, "n2: ",n2, "n3: ",n3)
	
	// 匿名变量跳过某些值
	const (
		m1 = iota // 0
		m2        // 1
		_         // 2
		m4        // 3
	)
	fmt.Println("m1: ",m1, "m2: ",m2, "m4: ",m4)

	// iota声明中间插队
	const (
		j1 = iota // 0
		j2 = 100  // 100
		j3 = iota // 2
		j4        // 3
	)
	const j5 = iota   // 0

	fmt.Println("j1: ",j1, "j2: ",j2, "j3: ",j3, "j4: ",j4, "j5: ",j5)


        const (
		A = iota    // 0
		B           // 1
		C           // 2
		D = "haha"  // haha   iota = 3
		E           // haha   iota = 4
		F = 100     // 100    iota = 5
		G           // 100    iota = 6
		H = iota    // iota = 7
		I           // iota = 8
	)
	const (
		J = iota    // 0
	)
        
	fmt.Println(A, B, C, D, E, F, G, H, I, J)


	// 定义数量级.使用 iota 结合 位运算 表示资源状态的使用案例
	const (
		Open = 1 << iota  // 0001
		Close             // 0010
		Pending           // 0100
	)
	
	const (
		_  = iota		// 使用 _ 忽略不需要的 iota
		KB = 1 << (10 * iota)	// 1 << (10*1)
		MB = 1 << (10 * iota)	// 1 << (10*2)
		GB = 1 << (10 * iota)	// 1 << (10*3)
		TB = 1 << (10 * iota)	// 1 << (10*4)
		PB = 1 << (10 * iota)	// 1 << (10*5)
	)
	fmt.Println("KB: ",KB, "MB:",MB, "GB:",GB, "TB: ",TB, "PB: ",PB)


	// 多个iota定义在一行.赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
	const (
		Apple, Banana = iota + 1, iota + 2 // Apple=1 Banana=2
		Cherimoya, Durian                  // Cherimoya=2 Durian=3
		Elderberry, Fig                    // Elderberry=3, Fig=4
	)
	fmt.Println("Apple: ",Apple, "Banana: ",Banana, "Cherimoya:",Cherimoya, "Durian: ",Durian, "Elderberry:",Elderberry, "Fig: ",Fig)

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
