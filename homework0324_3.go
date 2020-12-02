package main

import "fmt"

func main() {
	fmt.Println("请输入一个整数：")
	var num int
	fmt.Scanln(&num)
	year(num)

}
func year(n int)  {
	if n%4 == 0 && n%100 != 0 {
		fmt.Println("是闰年")
	}else {
		fmt.Println("不是闰年")
	}
}