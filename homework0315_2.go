package main

import "fmt"

func main() {
	numArr :=[]int{5,9,20,120,450,340}
	min :=0
	for index := 0; index < len(numArr); index++ {
		min = numArr[index]
		for i := index + 1; i < len(numArr); i++ {
			if numArr[i]>min {
				numArr[index],numArr[i]=numArr[i],numArr[index]
				min = numArr[index]
			}
		}
	}
	fmt.Println(numArr)
}
