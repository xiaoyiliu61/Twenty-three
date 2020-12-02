package main

import (
	"fmt"
	"strings"
)

func main() {
 str:="http://192.168.10.100:8080/Day33_Servlet/aa.jpeg"
 str1:=str[27:]
 s:=strings.SplitAfterN(str1,".",1)
 //fmt.Println(s)
 count:=0
 m:=make(map[string]int)
	for _,value:=range s{
		s1Slice:=strings.Split(value,".")
		fmt.Println(s1Slice[0])
		count++
		m[s1Slice[1]]++
	}
	fmt.Print("文件类型是：")
	for key,_:=range m{
		fmt.Print(key+"\t")
	}
}
