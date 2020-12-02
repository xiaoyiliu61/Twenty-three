package main

import "fmt"

func main() {
	var n =6
	for i := 1;i<=n ;i++  {
		for j := 0; j<n-i;j++  {
			fmt.Print(" ")
		}
		 fmt.Print("*********")
		fmt.Printf("\n")
		}
	} 	

