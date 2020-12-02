package main

import "fmt"

func main() {
	var n =5
	for i:=1;i<=n;i++{
		for j:=0;j<n - i;j++{
			fmt.Print(" ")
		}
		for m:=0;m<2*i-1;m++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i:=1;i<n;i++{
		for j:=0;j<i;j++{
			fmt.Print(" ")
		}
		for m:=0;m< 2*n-1 - 2*i;m++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}