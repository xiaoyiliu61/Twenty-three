package main

import "fmt"

func main() {
	arr:=[]int{97,98,99,100}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		fmt.Printf("%c\n",arr[i])
	}
}
