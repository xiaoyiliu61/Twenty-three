package main

import (
	"fmt"
	"math"
)

func main() {
	r1:=2.3
	s:=area(r1)
	fmt.Printf("圆形的半径是：%f,圆形的面积是%.3f\n",r1,s)
    r2:=5.1
    s2,l2:=Length(r2)
    fmt.Println(s2,l2)
}
func area(n float64) float64{
	area:=math.Pi*n*n
    return area
}
func Length(r float64)(float64,float64)  {
	area:=math.Pi*r*r
	length:=math.Pi*r*2
	return area,length
}

