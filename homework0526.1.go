package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connStr := "root:409216@tcp(127.0.0.1:3306)/liuyi?charset=utf8"
	database, err := sql.Open("mysql",connStr)
	defer database.Close()
	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("连接失败")
		return
	}
	fmt.Println("连接成功")

	/**
	  向数据库表中插入一条数据
	   1、向哪个表中插入：stu_info
	    2、插入哪些字段: name, sex
	    3、插入的具体的数据: 刘翔，男
	*/
	//prepare：准备。准备执行sql语句
	stmt, err := database.Prepare("insert into stu_info(id ,name, sex ) values(1, '刘义', '男')")
	stmt, err = database.Prepare("insert into stu_info(id, name, sex ) values(2, '胡子涵', '男')")
	if err != nil{
		fmt.Println("插入失败")
		return
	}
	//execute:执行   window：.exe
	//真正执行sql语句，将数据操作更新到数据库当中
	result, err := stmt.Exec()
	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("插入数据执行失败")
		return
	}

	//查看sql操作后的结果: sql语句的操作影响了几行数据
	rowsAffected, err :=result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("影响了%d行\n",rowsAffected)
}
