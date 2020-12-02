package main

import "fmt"
//定义一个长度为5的数组，并初始化数组的元素值，定义一个p，
//并赋值为素组的地址。要求：打印出数组的地址（十六进制的格式），并打印出变量p的类型。
func main() {
	arr:=[5]int{}
    fmt.Printf("数组arr的地址：%p\n",&arr)
	p:=&arr
	fmt.Printf("变量p的类型：%T\n",p)
}
