package main

import "fmt"

func main() {
	var n =4
	for a := 1; a<=n;a++  {
		for b:=n;b>=a ;b--  {
			fmt.Print(" ")
		}
		for c := 1; c <= 2*a-1; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
