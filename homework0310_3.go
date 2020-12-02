package main

import "fmt"

func main() {
	numArr :=[4]float64{2.23,4.32,6.92,8.58}
	fmt.Println(numArr)
	sum :=float64(0)
	for _,value:=range numArr{
		sum += value
	}
	fmt.Printf("元素相加之和是：%f\n",sum)
}