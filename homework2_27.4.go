package main

import "fmt"

func main() {
	a := 4
	b := 3
	res1 := a < b && b / 2 == 1 && a % 3 != 0
	res2 := (a+b)*3 < a<<2 || (a-b) >0
	fmt.Println(res1)
	fmt.Println(res2)
}
