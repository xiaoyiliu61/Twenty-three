package main

import "fmt"

func main() {
	var num int
	fmt.Println("给定数字为：")
	fmt.Scanln(&num)

	if num==1{
		fmt.Println("今天星期一")
	}else if num==2 {
		fmt.Println("今天星期二")
	}else if num==3{
		fmt.Println("今天星期三")
	}else if num==4{
		fmt.Println("今天星期四")
	}else if num==5{
		fmt.Println("今天星期五")
	}else if num==6{
		fmt.Println("今天星期六")
	}else if num==7{
		fmt.Println("今天星期日")
	}else{
		fmt.Println("信息错误")
	}
}
