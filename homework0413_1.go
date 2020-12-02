package main

import (
	"fmt"
	"math"
)

//定义圆形结构体，包含有半径、周长、面积等属性；
//定义圆形结构体所拥有的两个方法：
//计算面积、计算周长。在main函数中实现调用。
/*func main() {
    circular:=Circular{
		r: 3,
	}
	p1:=Area1(circular.r)
	fmt.Println("圆形的面积：",p1)
    p2:=Perimeter1(circular.r)
    fmt.Println("圆形的周长：",p2)

}

type Circular struct {
	r,c,s float64
}

func Area1(r1 float64) float64 {
	return r1*r1*math.Pi
}
func Perimeter1(r2 float64)float64  {
	return 2*math.Pi*r2
}*/
func main() {
	circle:=Circle{
		r: 2.1,
	}
	circle.s=circle.calArea()
	fmt.Println(circle.s)

	circle.l=circle.calLength()
	fmt.Println(circle.l)
}

type Circle struct {
	r float64
	l float64
	s float64
}

func (c Circle)calArea() float64  {
	c.s=math.Pi*c.r*c.r
	return  c.s
}

func (c Circle) calLength() float64  {
	c.l=math.Pi*2*c.r
	return c.l
}

