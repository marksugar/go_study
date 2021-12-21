package main

import "fmt"

/*
值传递：将原来的值拷贝了一份，传递给变量(或者不一定是变量，可以是数组等),传递的都是数据的副本(拷贝).
	如：int,float,string,bool,array
引用传递：引用的是内存地址，而不是拷贝原先的值.
	如：slice、map
*/
func main() {
	num := 10
	fmt.Printf("%T\n", num) // int

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [3]float64{2.15, 3.18}
	arr3 := [4]int{5, 6, 7, 8}
	arr4 := [4]string{"hello", "allenjol"}

	fmt.Printf("%T\n", arr1) // [4]int
	fmt.Printf("%T\n", arr2) // [3]float64
	fmt.Printf("%T\n", arr3) // [4]int
	fmt.Printf("%T\n", arr4) // [4]string

	// 赋值
	num2 := num            // 值传递.就是将变量num中的值拷贝了一份，传递给了num2变量
	fmt.Println(num, num2) // 10 10

	num2 = 20
	fmt.Println(num, num2) // 10 20

	// 数组
	arr5 := arr1      // 值传递。将arr1数组中的值拷贝了一份给数组arr5
	fmt.Println(arr1) // [1 2 3 4]
	fmt.Println(arr5) // [100 2 3 4]

	arr5[0] = 100 // 将arr5数组中的索引为0的值更改为100，不会影响数组arr1的值。这叫做值传递
	fmt.Println(arr1)
	fmt.Println(arr5)

	// 值类型还可以比较值是否相等。
	a := 3
	b := 4
	fmt.Println(a == b)       // 比较a和b的数值是否相等
	fmt.Println(arr5 == arr1) // 数组会比较每个元素的类型，长度都要一致，然后对比对应的位置上的元素是否相等
	//fmt.Println(arr2 == arr3)	// 报错，类型不匹配
}
