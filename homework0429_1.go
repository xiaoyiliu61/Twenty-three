package main

import (
	"fmt"
	"net/http"
)

//编程实现一个web服务器，可以处理名为/login的请求，
//模拟登陆请求，携带用户名和密码两个数据。
//在服务器端接收登陆的数据，并对请求数据进行判断。
//如果用户名是hello，密码是123456，表示用户存在，
//返回“恭喜登录成功“结果字样，否则返回”对不起，您的账号或者密码不正确，请重新尝试“的提示语。
func main() {
  /*  http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("用户名：\n 密码:"))

		url:=request.URL
		fmt.Println("path:",url.Path)
		fmt.Println("请求方法：",request.Method)

		name:=request.FormValue("name")
		fmt.Println(name)
		password:=request.FormValue("password")
		fmt.Println(password)

		if name=="hello"&&password=="123456" {
			fmt.Println("恭喜登录成功")
		}else {
			fmt.Println("你的账户和密码不正确，请重新尝试")
		}
	})

    fmt.Println("启动监听服务")

    err:=http.ListenAndServe("127.0.0.3:8080",nil)
	if err!=nil {
		fmt.Println(err.Error())
	}

   */
	http.HandleFunc("/login",login)

	err:=http.ListenAndServe("127.0.0.1:8080",nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
func login(writer http.ResponseWriter, request *http.Request) {
	username:=request.Form.Get("username")
    password:=request.Form.Get("password")
	if username=="hello"&&password=="123456" {
		writer.Write([]byte("登录成功"))
	}else {
		writer.Write([]byte("密码错误"))
	}
}
