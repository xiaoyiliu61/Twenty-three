package main

import (
	"fmt"
	"net/http"
)

//作业一：
//① 编写一个web服务器程序，在自己电脑上的9001端口进行监听，
//可以处理/userlogin请求,返回一句话“欢迎访问用户登录功能”
//② 编程实现一个客户端程序，请求步骤①中的服务器，并访问/userlogin，
//在客户端程序中接收请求，并打印出请求到的信息。
//程序编写好后，先运行服务器程序，然后运行客户端程序，查看程序运行效果。
/*func main() {
	client:=http.Client{}

	http.HandleFunc("/userlogin",UserLogin)

	err:=http.ListenAndServe("127.0.0.1:9001",nil)
	if err != nil {
		fmt.Println(err.Error())
	}

    urlStr:="http://127.0.0.1:9001/userlogin"
    params:=&url.Values{
    	"inform":{"欢迎访问用户登录功能"},
	}
	bufferstring:=bytes.NewBufferString(params.Encode())
	response,err:=client.Post(urlStr,"application/x-www-form-urlencoded",bufferstring)
	if err != nil {
			log.Fatal(err)
			return
		}
	   defer response.Body.Close()

	fmt.Println(response.StatusCode)
}
func UserLogin( writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("欢迎访问用户登录功能"))
}*/
func main() {
	http.HandleFunc("/userlogin", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("用户访问登录功能"))
	})

	err:=http.ListenAndServe("127.0.0.1:9001",nil)
	if err!=nil {
		fmt.Println(err)
	}
}