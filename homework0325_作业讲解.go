package main

import (
	"fmt"
	"strings"
)

func main() {
	num1:=5
	num2:=6
	length(num1,num2)
	st:="WUAHNJIAYOUZHONGGUOJIAYOU"
	s:=strCount(st)
	fmt.Println(s)
}
func length(a int,b int)int{
	l:=(a+b)*2
	return l
}
func strCount(str string) map[string]int {
	m:=make(map[string]int)
	s:=""
	for index := 0; index < len(str); index++ {
		s=string(str[index])
		if m[s]==0 {
			m[s]=strings.Count(str,s)
		}
	}
	return m
}
