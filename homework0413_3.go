package main

import "fmt"

//使用接口定义汽车的标准：① 汽车品牌  ② 汽车能够驱动
//并定义结构体，实现卡车，电动汽车、三轮车三种定义，实现汽车接口标准。
//分别实例化解放牌卡车，特斯拉电动车，时风三轮车，并调用对应的驱动方法
/*func main() {
   c:=Car{
	   kind: "卡车",
	   c_brand: "解放牌卡车",
   }
  c.LinkIn()

   c1:=Car{
	   kind:   "电动汽车",
	   c_brand: "特斯拉电动车",
   }
   c1.LinkIn()

   c2:=Emplex{
		Car:Car{
			kind:   "三轮车",
			c_brand: "时风三轮车",
		},
		price: 1000,
		color: "红色",
	}
	fmt.Printf("种类：%s,品牌：%s,价格：%d,颜色：%s",c2.kind,c2.c_brand,c2.price,c2.color)
}


func TestCar(car2 CarStandard){
	car2.LinkIn()
}

type CarStandard interface {
	Brand() string
	LinkIn ()
}

type Car struct {
	c_brand string
	kind string
}

func (c Car) name() string  {
	return c.c_brand
}
func (c Car) LinkIn()  {
	fmt.Println("车的种类：",c.kind,"车的品牌：",c.c_brand)
}

type Emplex struct {
	Car
	price int
	color string
}*/
func main() {
	jiefang:=Truck{
		t_name: "解放",
		weight: 50,
	}
	jiefang.drive()

	telsa:=EVE{
		e_name: "特斯拉",
		power:  350,
	}
	carDiver(telsa)

	shifeng:=Sanlun{
		s_name: "时风",
		num:    3,
	}
	shifeng.drive()

	var c  Car
	c=Sanlun{
		s_name: "野马三轮",
		num:    3,
	}
	c.drive()

	fmt.Printf("%T\n",c)

}

func carDiver(car Car)  {
	car.drive()
}

type Car interface {
	name() string
	drive ()
}

type Truck struct {
	t_name string
	weight int
}

func (t Truck) name() string  {
	return t.t_name
}

func (t Truck) drive()  {
	fmt.Printf("%s牌卡车载重是%d吨\n",t.t_name,t.weight)
}

type EVE struct {
	e_name string
	power int
}

func (e EVE) name() string {
	return e.e_name
}

func (e EVE) drive()  {
	fmt.Printf("%s牌电动车电池续航里程是%dKM\n",e.e_name,e.power)
}

type Sanlun struct {
	s_name string
	num int
}

func (s Sanlun) name() string  {
	return s.s_name
}

func (s Sanlun) drive()  {
	fmt.Printf("%s牌车有%d个轮子\n",s.s_name,s.num)
}