package main

import "fmt"

//定义一个函数，实现两数相加，要求最后返回一个指针类型，并在main中调用和接收，最后打印计算后的结果值
/*func main() {
	a:=6
	b:=1
	r1:=zhizheng(a,b)
	fmt.Println("两数相加的值：",*r1)

}
func zhizheng(a,b int) *int {
	c:=a+b
	var p *int
	p=&c
	return p
}*/
func main() {
	num3,num4:=45,12
	r1:=add(num3,num4)
	fmt.Println(*r1)
}

func add(a int, b int) *int {
	s:=a+b
	return &s
}
