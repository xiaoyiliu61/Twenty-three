package main

import (
	"fmt"
	"log"
	"net/http"
)

//作业一：
//① 编写user_info.html，该页面用于实现用户的实名信息认证。
//该页面上包含的信息有：真实姓名，性别（单选框），出生年月，家庭住址，身份证号，
//学历（下拉选择框），发证机关（xxx公安局），该页面有两提交按钮，点击提交按钮，可以提交数据到服务器。
//②使用Go语言程序编写一个服务器程序，在本地机上运行，接收user_info页面提交的数据。
//在服务器端解析前端提交的数据，并对数据做检查，最后根据检查返回提示信息。
//
//数据规范：
//姓名：最长不能超过12个字，名字不能为空，否则提示姓名不符合规范。
//性别：有男、女两个选项，两个选项必须选择一个，否则提示：请选择性别信息。
//出生年月：出生年月日不能超过2020年5月9日，否则提示：太着急了，还没到日子呢。
//身份证号：长度必须为18位，否则提示身份证长度不符合规范。
func main() {
   http.HandleFunc("/user_info",UserInfo1)

   err:=http.ListenAndServe("127.0.0.1:8008",nil)
	if err != nil {
		log.Fatal(err)
	}
}
func UserInfo1(writer http.ResponseWriter,request *http.Request)  {
	 //fmt.Println("实现用户的实名信息认证")
	 request.ParseForm()
	 user:=request.Form.Get("username")
	if len(user) > 12 && len(user) == 0 {
		fmt.Fprintf(writer,"姓名不符合规范")
	}
	 sex:=request.Form.Get("sex")
	 age:=request.Form.Get("age")
	 pwd:=request.Form.Get("pwd")
	if len(pwd) != 18 {
		fmt.Fprintf(writer,"身份证长度不符合规范")
	}
	 fmt.Println(user)
	 fmt.Println(sex)
	 fmt.Println(age)
	 fmt.Println(pwd)

}