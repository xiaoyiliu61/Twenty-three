package main

import "fmt"

func main() {
  var usename string
  var password string
  fmt.Println("请输入你的名字： ")
  fmt.Println("请输入你的密码： ")
  fmt.Scanln(&usename,&password)
  fmt.Printf("用户名：%s\n密码：%s",usename,password)
}