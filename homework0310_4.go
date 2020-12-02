package main

import "fmt"

func main() {
	var numArr [10]int
    fmt.Println(numArr)
	for index := 0;index<=len(numArr)-1 ;index++  {
		fmt.Printf("请输入第%d个元素数值：\n",index+1)
		num:=0 //fmt.scanlm(&numarr[index])
		fmt.Scanln(&num)
		numArr[index]=num
	}
	for index, value := range numArr{//for_,value:=range numarr
		fmt.Println(index,value)//fmt.printf("%d\t",value)
	}

}
