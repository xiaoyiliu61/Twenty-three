package main

import "fmt"

/*func main() {
	arr:=[4]int{4,9,1,6}
	var p1 *[4]int
	p1=&arr
	for i := 1; i < len(p1); i++ {
		for index := 0; index < len(p1)-1; index++ {
			if p1[index]>p1[index+1] {
				p1[index],p1[index+1]=p1[index+1],p1[index]
			}
		}
	}
	fmt.Println(*p1)

}*/
func main() {
	arr:=[3]int{3,5,1}
	p1:=&arr
	fmt.Printf("%T\n",p1)

	num:=100
	p2:=&num
	fmt.Printf("%T\n",p2)

	num1:=1
	num2:=2
	num3:=3

	var arr1 [3]*int
	arr1[0]=&num1
	arr1[1]=&num2
	arr1[2]=&num3
	fmt.Println(arr1)

	fmt.Println("========")
	arr2:=[4]int{1,2,3,4}
	fmt.Printf("数组的地址：%p\n",&arr2)
	p3:=&arr2
	fmt.Printf("%p\n",p3)
	fmt.Println(&arr2[0])

	fmt.Println("=======")
	var p *[4]int
	p=&arr2
	fmt.Printf("%p\n",p)

	a:=10
	b:=12
	c:=15
	arr3:=[3]*int{&a,&b,&c}
	fmt.Println(arr3)

	arr4:=[5]int{12,23,32,5,9}
	s1:=sort(&arr4)
	fmt.Println(*s1)
}
func sort(p1 *[5]int)*[5]int  {
	for i := 1; i < len(p1); i++ {
		for index := 0; index < len(p1)-1; index++ {
			if p1[index]>p1[index+1] {
				p1[index],p1[index+1]=p1[index+1],p1[index]
			}
		}
	}
	return p1
}

