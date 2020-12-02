package main

import "fmt"

//定义一个长度为6的字符串数组，并进行元素的初始化赋值；
//要求：将字符串数组中各元素的地址保存到一个指针数组中，并打印该指针数组的元素值。
//最后通过操作指针数组元素修改原字符串数组中最后一个元素的元素值
/*func main() {
	arr1:=[6]int{1,2,3,4,5,6}
	r1:=zhizheng(&arr1)
	fmt.Println(*r1)

}
func zhizheng(p1 *[6]int)*[6]int  {
	for index := 0; index < len(p1); index++ {
		p:=&p1[index]
		fmt.Println(p)
	}
	p1[5]=61
	return p1
}*/
func main() {
	var num int =100
	fmt.Println(num)
	num1:=101
	fmt.Println(num1)
	pStr1:=[6]*string{}
	arrNum1:=[6]string{"wu","han","jia","you","jia","you"}
	for index := 0; index < len(arrNum1); index++ {
		pStr1[index]=&arrNum1[index]
	}
	fmt.Println(pStr1)
	fmt.Println(pStr1[5])
	fmt.Println(*pStr1[5])
	*pStr1[5]="jiayou"
}
