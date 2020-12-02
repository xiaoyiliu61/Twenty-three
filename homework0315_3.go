package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	numArr :=[10]int{}
	rand.Seed(time.Now().UnixNano())
    num:=0
    count:=0
	for i:= 0; i< len(numArr); i++ {
		OUT:
			num=rand.Intn(16)
		count++
		for index := 0; index < i; index++ {
			if num==numArr[index]{
			 goto OUT
			}
		}
		numArr[i]=num
	}
	fmt.Println("最后的数组是：",numArr)
    fmt.Println("生成了%d次的随机数\n",count)
}
