package main

import "fmt"

func main() {
  var num int
  fmt.Println("给定一个数字：")
  fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("今天星期一")
	case 2:
		fmt.Println("今天星期二")
	case 3:
		fmt.Println("今天星期三")
	case 4:
		fmt.Println("今天星期四")
	case 5:
		fmt.Println("今天星期五")
	case 6:
		fmt.Println("今天星期六")
	case 7:
		fmt.Println("今天星期七")
	default:
		fmt.Println("信息错误")
	}
}
