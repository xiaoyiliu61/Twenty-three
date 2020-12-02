package main

import "fmt"

func main() {
	var num int
	for num = 100; num < 1000;num++  {
		if num<1000{
		 a :=num%10
		 b :=num/10%10
		 c :=num/100
			if  num==a*a*a+b*b*b+c*c*c {
				fmt.Println(num)
			}

		}
	}
}
