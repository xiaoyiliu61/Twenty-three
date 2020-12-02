package main

import (
	"fmt"
	"strings"
)

func main() {
	/* var num =[]string{"WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"}
    num=strings.Split("WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA","")
	var num2 =make(map[string]int)
		for  _,value:=range num{
			if num2[value]!=0 {
				num2[value]++
			}else {
				num2[value]=1
			}
	}
	fmt.Println(num2)*/
	//给定字符串，打印出字符串当中每个字母出现的次数
	str:="WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	fmt.Println(str)
    //方法一
    //str[x:y]通过下标获取某一段字符串
    s:=""
    m:=make(map[string]int)//string表示每个字母，int表示字母出现的个数
    //map的特点：key唯一，value值可以修改
	for index := 0; index < len(str); index++ {
		s=str[index:index+1]
		if m[s]>0 {//如果m[s]>0，则表示map当中已经有该字母的次数统计了
			continue
		}
		m[s]=1//先赋值为1m["w"]=1 m["u"]=1
		for i := index + 1; i < len(str); i++ {//从遇到的字母的下一位开始，遍历到字符串的末尾，统计字母出现的个数n-1
			if str[i:i+1] == s {//出现了一个相同的字母，计数要+1
				m[s]++
			}
		}
		fmt.Println(m)
	}


	//方法二
	str1:="WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	// 数组
	//1、将字符串变成一个数组 [35]string{"W","U","H","A","N","J",.....}
	//2、利用map数据结构本身的特点进行统计
	//   map数据结构：1、key唯一、与value对应
	//               2、map[key] = value1、value2 如果key已经存在，value会修改成新的值
	//               3、如果不存在key值，则map[key]返回一个默认值 int类型的默认值是0
    strArr:=[35]string{}
	for index := 0; index < 35; index++ {//获取到每一个字母，放到数组当中
		strArr[index]=str1[index:index+1]
	}
	map1:=make(map[string]int)
	count:=0//变量用于获取该字母的次数
	for index := 0; index < len(strArr); index++ {//拿到数组中的每个元素
		key:=strArr[index]
		count=map1[key]
		count++//对count值进行++，增加1
		map1[key]=count
	}
	//遍历输出
	for key,value:=range map1{
		fmt.Println(key,value)
	}

	//第三种
	//第三种解决方案
	//对字符串转数组，进行优化
	// "YU HONG WEI" 按照空格进行切割
	//字符串转切片的方法: 字符串切割
	str2:= "WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	str3:=strings.Split(str2,"")

	map2:=make(map[string]int)//value=map[key]
	for index := 0; index < len(str3); index++ {
		map2[str3[index]]++
	}
	fmt.Println(map2)


}
