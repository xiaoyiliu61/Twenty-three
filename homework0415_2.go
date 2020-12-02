package main

import (
	"fmt"
)

//定义长度为6的数组并进行赋初始值。
//使用for循环访问0到10的数组下标上的元素。
//处理程序可能遇到的异常，处理方式主动终止程序执行，并给出异常的原因
func main() {
	arr:=[6]string{"a","b","c","d","e","f"}
	for index := 0; index<10; index++ {
		if index < len(arr) {
			fmt.Println(arr[index])
		} else {
			panic("程序错误")
		}
	}
}