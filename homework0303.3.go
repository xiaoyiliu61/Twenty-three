package main

import "fmt"

func main() {
	fmt.Println("请输入两个整数：")
	var num1 int
	var num2 int
	fmt.Scanln(&num1,&num2)

	fmt.Println("请输入要执行的操作符号（+ — * / ++ -- %):")
	var str string
	fmt.Scanln(&str)

	switch str {
	case "+":
		fmt.Println("两数相加的结果是：",num1+num2)
	case "-":
		fmt.Println("两数相减的结果：",num1-num2)
	case "*":
		fmt.Println("两数相乘的结果是：",num1*num2)
	case "/":
		fmt.Println("两数相除的结果：",num1/num2)
	case "++":
		num1++
		num2++
		fmt.Println("num1和num2加一的值",num1,num2)
	case "--":
		num1++
		num2++
		fmt.Println("num1和num2减一的值",num1,num2)
	case "%":
		fmt.Println("num1%num2的值：",num1%num2)
	}
}
