package main

import (
	"fmt"
	"net/http"
)

//编写一个web服务，在浏览器中访问该服务时，
//显示自己硬盘上某个存在的文目录。服务的端口是9000。
func main() {
	err:=http.ListenAndServe("127.0.0.1:9000",http.FileServer(http.Dir("C:/Users/pc/GoProjects/src/HomeHORK")))
	if err != nil {
		fmt.Println(err.Error())
	}
}
