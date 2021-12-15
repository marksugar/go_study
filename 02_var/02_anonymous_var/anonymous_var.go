package main

import "fmt"

func GetData() (int, string) {
	return 100, "allen"
}

func main() {
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}
