package main

import "fmt"

func main() {
	/*var l =[20]int{3,9,10,11,8,4,9,1,20,10,11,21,19,3,8,4,1,10,20,21}
	var y =make(map[int]int)
	for _,value:=range l {
		if y[value]!=0 {
			y[value]++
		}else {
			y[value]=1
		}
	}
	fmt.Println(y)*/
	numArr := [20]int{3,9,10,11,8,4,9,1,20,10,11,21,19,3,8,4,1,10,20,21}
	map1:=make(map[int]int)
	for index := 0; index < len(numArr); index++ {
		map1[numArr[index]]++
	}
	fmt.Println(map1)
}
