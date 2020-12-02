package main

import "fmt"

func main() {
	fmt.Println("请输入两个数")
	var a ,b int
	fmt.Scanln(&a,&b)
	c:=Perimeter(a,b)
	fmt.Println("矩形周长是：",c)
	Perimeter(a,b)
}
func Perimeter(num1 int,num2 int) int  {
	a:=(num1+num2)*2
	return  a
}
