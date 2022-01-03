package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	//因为ParseBool会返回两个值，value和error
	//因为我只需要value，不需要error，所以用_忽略返回的error值
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T, value b=%v\n", b, b)

	var str2 string = "123450"
	var n1 int64
	var n2 int
	// 因为上面返货的n1是int64，所以这里要用int64来接收，因此n1的类型为float64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	// 因为上面返回的n1是int64，这里想要得到int类型，所以要强制转换为int类型
	n2 = int(n1)
	fmt.Printf("n1 type %T, value n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T, value n2=%v\n", n2, n2)

	var str3 string = "123.56"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T, value f1=%v\n", f1, f1)
}
