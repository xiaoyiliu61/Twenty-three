package main

import (
	"errors"
	"fmt"
)

//有结构体Person，包含name、age、address等属性。
//定义一个函数接收结构体类型，判断传入的person的年龄是否已经成年
//，如果已成年，返回true，如果未成年，返回false；
//如果年龄是0或者负数，使用error返回年龄不合法等提示信息
func main() {
	a:=Person{
		name:    "刘义",
		age:     18,
		address: "江西省九江市",
	}
	 isYoung,err:=a.CheckAge()
	if err!=nil {
		fmt.Println(err)
		return
	}
	if isYoung {
		fmt.Println("已成年")
	}else {
		fmt.Println("未成年")
	}
}

type Person struct {
	name string
	age int
	address string
}

func (per Person)CheckAge() (bool,error) {

	if per.age<18 {
		 return false,nil
	}else if per.age>=18 {
	    return true,nil
    }else if per.age<=0 {
		return false,errors.New("年龄不合法")
	}
	return false,nil
}