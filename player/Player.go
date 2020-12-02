package player

import "fmt"

type Athletes struct {
	Name string
	Sex string
	Grade int
	Project string
}

func (a Athletes) Run() {
	fmt.Printf("姓名：%s,性别：%s,成绩：%d,项目：%s",a.Name,a.Sex,a.Grade,a.Project)
}