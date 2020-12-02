package main

import "fmt"

func main() {
	var n = 7
	for i := 1; i <n; i++ {
		for k := 1; k <= 2*i-1; k++ {
			if i == 1||i==6 {
				fmt.Print("*")
			}else if k == 1 || k == 2*i-1 {
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

