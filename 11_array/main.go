package main

import (
	"fmt"
	"runtime"
)

func init() {
	var pro = "Enter a digit"
	if runtime.GOOS == "windows" {
		pro = fmt.Sprintf(pro, "ctrl + D, enter")
		fmt.Println("windows")
	} else {
		pro = fmt.Sprintf(pro, "command + D")
		fmt.Println("not windows")
	}

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("it's main function...")
	aabs := Abs(5)
	fmt.Println(aabs)
}
