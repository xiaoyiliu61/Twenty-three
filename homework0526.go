package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connStr:="root:409216@tcp(127.0.0.1:3306)/liuyi?charset=utf8"
	db,err:=sql.Open("mysql",connStr)
	if err!=nil{
		fmt.Println("数据连接失败",db)
		return
	}
	fmt.Println("连接成功")
	db.Close()

}
