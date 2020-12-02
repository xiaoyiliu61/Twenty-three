package main

import "fmt"

func main() {
	numArr :=[]int{5,9,20,120,450,340}

	for i := 1; i < len(numArr); i++ {
		for index := 0; index < len(numArr)-1; index++ {
			if numArr[index]<numArr[index+1] {
				numArr[index],numArr[index+1]=numArr[index+1],numArr[index]
	     	}
		}
	}
	fmt.Println(numArr)
}