package main

import (
	"fmt"
)

func main() {
	num:=6
	r:=recursion(num)
	fmt.Println(r)
}
func recursion(n int) int {

	if n==1 {
		return 1
	}else {
      return n+recursion(n-1)*10
	}
}