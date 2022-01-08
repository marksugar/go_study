package main

import "fmt"

/*
	赋值运算符就是将某个运算后的值赋值给指定的变量

	c += a --> c = c + a
	c -= a --> c = c - a
	c *= a --> c = c * a
	c /= a --> c = c / a
	c %= a --> c = c % a

	<<= 左移后 再赋值	c <<= 2 --> c = c << 2
	>>= 右移后 在赋值 c >>= 2 --> c = c >> 2
	&= 按位与 后赋值  c &= 2 --> c = c & 2
	^= 按位异或 后赋值 c ^= 2 --> c = c ^ 2
	|= 按位或 后赋值   c |= 2 --> c = c | 2
*/

func test() int {
	return 90
}

func assign_operator() {
	var i int
	i = 10
	fmt.Println("i=", i)

	// 面试题：有两个变量a和b，不能用临时变量，怎么交换变量的值。如下：
	a := 9
	b := 2
	fmt.Printf("方法一：a和b交换前的值：a=%d,b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("方法一：a和b交换后的值：a=%d,b=%d\n", a, b)

	// 使用中间变量t来交换变量的值
	c := 10
	d := 20
	fmt.Printf("方法二：c和d交换前的值：c=%v,d=%v\n", c, d)
	t := c
	c = d
	d = t
	fmt.Printf("方法二：c和d交换后的值：c=%v,d=%v\n", c, d)

	// 复合赋值
	i += 7
	fmt.Println("i+7后的值：", i)

	// 赋值运算执行顺序是从右向左
	var e int
	e = a + 3
	fmt.Println("e=", e)

	// 赋值运算符的左边 只能是变量，右边可以使变量、敞亮、表达式
	var f int
	f = a // 变量
	fmt.Println("f=", f)
	f = 8 + 2*8 // = 右边是表达式
	fmt.Println("f=", f)
	f = test() + 90 // = 右边是表达式
	fmt.Println("f=", f)
	f = 890 // 890 是常量
	fmt.Println("f=", f)

	// 面试题：两个变量:k,l，要求将其进行交换，不能用中间变量，最终打印结果
	var k int = 10
	var l int = 20

	k = k + l
	l = k - l // l = k + l - l ---> k = l
	k = k - l // k = k + l - k ---> l = k
	fmt.Printf("k=%v,l=%v\n", k, l)
}

func main() {
	assign_operator()
}
