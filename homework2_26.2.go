package main

import "fmt"

func main() {
	people :="wuhanjiayou,zhongguojiayou"
	fmt.Println("字符串people的长度",len(people))
	fmt.Println(people[:5])
	fmt.Println(people[12:20])
	fmt.Println(people[:11])
	fmt.Println(people[12:])




}
