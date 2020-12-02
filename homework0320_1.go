package main

import "fmt"

func main() {
   str:="我爱中国，I LOVE CHINA"
   fmt.Println(len(str))
   str1:=[]rune(str)
	for index := 0; index < len(str1); index++ {
		fmt.Printf("%c\n",str1[index])
	}
}
