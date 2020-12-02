package main

import "homehork/player"

//定义运动员结构体，运动员有姓名、性别、运动项目、成绩等属性；
//运动员有一个run方法，在run方法中打印运动员所属的运动项目。
//将运动员定义在包player中，在main包中实现调用
func main() {
  ath:=player.Athletes{
	  Name:    "刘义",
	  Sex:     "男",
	  Grade:   100,
	  Project: "跑步",
  }
  ath.Run()
}
