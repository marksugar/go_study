package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.7.md

	/*
		前缀和后缀，判断以什么开头、以什么结尾
		HasPrefix 判断字符串 s 是否以 prefix 开头：
			strings.HasPrefix(s, prefix string) bool
		HasSuffix 判断字符串 s 是否以 suffix 结尾
			strings.HasSuffix(s, suffix string) bool
	*/

	var str1 = "This is a pig?"
	fmt.Printf("Does the string \"%s\" have prefix \"%s\"\n", str1, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str1, "Th"))

	fmt.Printf("Does the string \"%s\" have suffix \"%s\"\n", str1, "pig?")
	fmt.Printf("%t\n", strings.HasSuffix(str1, "pig?"))

	fmt.Println("----------")

	/*
		字符串包含关系
		Contains 判断字符串 s 是否包含 substr
			strings.Contains(s, substr string) bool
	*/
	fmt.Printf("Does the string \"%s\" contains \"%s\"\n", str1, "pig")
	fmt.Printf("%t\n", strings.Contains(str1, "pig"))

	/*
		判断子字符串或字符在父字符串中出现的位置（索引）
			Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
				strings.Index(s, str string) int
			LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
				strings.LastIndex(s, str string) int
			查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位:
				strings.IndexRune(s string, r rune) int
	*/

	fmt.Println("----------")

	var str2 = "Hi, I'm Maric,Hi."
	fmt.Printf("The position of \"Maric\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "Maric"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "Hi"))

	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str2, "Hi"))

	fmt.Printf("Ths position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "Burger"))

	/*
		字符串替换
		Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new
			strings.Replace(str, old, new, n) string
	*/
	fmt.Println("----------")

	var str3 = "This phone is old,but it not old enough！"
	fmt.Println("Replace str: ", strings.Replace(str3, "old", "new", -1))
	fmt.Println("Replace str: ", strings.Replace(str3, "phone", "Girl", -1))

	/*
		统计字符串出现次数
			Count用于计算字符串str在字符串s中出现的非重叠次数
			strings.Count(s, str string) int
	*/
	fmt.Println("----------")

	var str4 = "Hello, how is it going, Hugo?"
	var manyG = "gggggggg"

	fmt.Printf("Number of H's in %s is: ", str4)
	fmt.Printf("%d\n", strings.Count(str4, "H"))

	fmt.Printf("Number of g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "g"))

	/*
		重复字符串
			Repeat 用于重复 count 次字符串 s 并返回一个新的字符串
			strings.Repeat(s, count int) string
	*/
	fmt.Println("----------")

	var origS = "Hi is me! "
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)

	/*
		修改字符串大小写
		ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符：
			strings.ToLower(s) string
		ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符：
			strings.ToUpper(s) string
	*/
	fmt.Println("----------")
	var orig = "Hey,how are your George?"
	var (
		lower string
		upper string
	)

	fmt.Printf("The original string is: %s\n", orig)

	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase strings is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase strings is: %s\n", upper)

	/*
		修剪字符串，去掉字符串空白
		使用 strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号；如果你想要剔除指定字符，
		则可以使用 strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉。
		该函数的第二个参数可以包含任何字符，如果你只想剔除开头或者结尾的字符串，
		则可以使用 TrimLeft 或者 TrimRight 来实现。
	*/
	fmt.Println("----------")
	var str5 = "  This is your borather..   ,right?   "
	fmt.Printf("Have not trimspace str is: %s\n", str5)
	fmt.Printf("trimspace str is: %s\n", strings.TrimSpace(str5))

	/*
		字符串分割
		strings.Fields(s) 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，
		并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。
		strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。
		因为这 2 个函数都会返回 slice，所以习惯使用 for-range 循环来对其进行处理
	*/

	/*
		拼接slice到字符串
		Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
			strings.Join(sl []string, sep string) string
	*/
	fmt.Println("----------")

	str6 := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str6)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	str7 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str7, "|")
	fmt.Printf("Splitted in slice: %s\n", sl2)
	for _, val :=range sl2{
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	str8 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by \";\" %s\n", str8)


	/*
	字符串和其他类型转换
	 */
	fmt.Println("----------")
	var orig2 = "666"
	var (
		an int
		newS2 string
	)

	// 注意64位系统和32位系统有区别
	fmt.Printf("The size if ints is: %d\n", strconv.IntSize)

	// 函数 strconv.Atoi 的作用是将一个字符串转换为一个整数
	an,_ = strconv.Atoi(orig2)
	fmt.Printf("The intger is: %d\n", an)
	an +=5
	newS2 = strconv.Itoa(an)
	fmt.Printf("The new strings is： %s\n", newS2)

}
