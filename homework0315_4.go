package main

import "fmt"

func main() {
	slice:=make([]int,5,7)
	slice[0]=1
	slice[1]=3
	slice[2]=5
	slice[3]=7
	slice[4]=9
	fmt.Println(slice)
	fmt.Println(len(slice),cap(slice))
}
