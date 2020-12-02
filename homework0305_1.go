package main

import "fmt"

func main() {
	var num int
	 fmt.Println("请输入一个数字")
	fmt.Scanln(&num)

	switch num {
	case 1 , 2, 3, 4, 5, 6, 7, 8, 9, 10:
		fmt.Println("上旬")
	case 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
		fmt.Println("中旬")
	case 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
		fmt.Println("下旬")

	}
}
