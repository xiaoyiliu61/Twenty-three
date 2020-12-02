package main

import "fmt"

/*
1.公司里的职员，可以描述职员的信息包括：员工编号，姓名，部门，薪水，交税数额。
 要求：编程创建结构体。并创建几个具体的职员信息，交税数额均设置为默认值0，
 然后打印输出这些职员的信息。
2.在作业一的基础上，创建一个函数，该函数接收职员类型，
在函数中判断职员的薪水是否大于5000，薪水小于等于5000元，交税数额为0；
如果大于5000，则计算员工要交的税额，并赋值到相应的字段。
最后在main函数中调用函数，并接受返回值，打印员工要交的税额。
注：员工应交税税额 = （薪水 – 5000） * 20 %
 */
func main() {
    figure1:=StaffInformation{
		StaffNumber:     0654503,
		StaffName:       "刘义",
		StaffDepartment: "成型加工",
		StaffSalary:     4500,
		StaffTollage:    0,
	}
	fmt.Println(figure1)

    var figure2 StaffInformation
    figure2.StaffName="何子笙"
    figure2.StaffDepartment="成型加工"
    figure2.StaffSalary=5500
    figure2.StaffTollage=0
    fmt.Println(figure2)

    r1:=Staff(figure1)
    r2:=Staff(figure2)
	fmt.Printf( figure1.StaffName+"应缴费:%d\n",r1)
	fmt.Printf( figure2.StaffName+"应缴费:%d\n",r2)

}
	type StaffInformation struct {
		StaffNumber int
		StaffName string
		StaffDepartment string
		StaffSalary int
		StaffTollage int

 }
 

func Staff(per StaffInformation) int {
     T:=0
	if per.StaffSalary > 5000{
		p:=(per.StaffSalary-5000)/5
        T=p
	}
	return T
}

