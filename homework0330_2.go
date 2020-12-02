package main

import "fmt"

func main() {
	func(n int){
		r:=1
		for index := n; index >=1; index-- {
			r*=index
		}
		fmt.Println(r)
	}(6)
}
