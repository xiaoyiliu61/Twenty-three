package main

import (
	"fmt"
	"math/rand"
	"time"
)

//在一个子协程中打印输出0-1000的随机整数，
//每次打印输出前，首先睡眠200ms，打印99个数字；
//在另一个子协程中，每间隔100ms输出一个一个大写字母（范围A-Z），
//打印199个字幕。在main中等待两个子协程任务执行完毕
func main() {

   go PrintNum()

   go PrintCode()

   time.Sleep(200*time.Second)
}

func PrintNum() {
	for index := 0; index < 99; index++ {
		fmt.Printf("序号：%d,数字%d\n",index,rand.Intn(1000))
		time.Sleep(200*time.Millisecond)
	}
}

func PrintCode() {
	for index := 0; index < 199;index++ {
		fmt.Printf("%d字符%s\n",index,string(65+rand.Intn(26)))
		time.Sleep(100*time.Millisecond)

	    fmt.Printf("%d字符%s\n",index,string(65+index%26))
	}
}