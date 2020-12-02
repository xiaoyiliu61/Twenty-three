package main

import "fmt"

func main() {
	var num int
	fmt.Println("给定一个月份（1-12）")
	fmt.Scanln(&num)
	switch num {
	case 3, 4, 5:
		fmt.Println("该月为春季")
	case 6, 7, 8:
		fmt.Println("该月为夏季")
	case 9, 10, 11:
		fmt.Println("该月为秋季")
	case 12, 1, 2:
		fmt.Println("该月为冬季")
	}

}
