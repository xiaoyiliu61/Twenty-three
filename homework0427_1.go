package main

import (
	"fmt"
	"math/rand"
	"time"
)


//启动一个goroutine，生成100个随机数发送到容量为5的缓冲通道ch1中，随机数间隔20ms生成一个；
//启动另外一个goroutine，从ch1中读取数据，然后计算读取到的数值的3次方，放到通道ch2中；
//在main函数中，从通道ch2中读取数据，并打印输出。最后程序结束。


func main() {
	ch1:=make(chan int,5)
	ch2:=make(chan int)

	go Goroutine1(ch1)
	go Goroutine2(ch1,ch2)


	for result:=range ch2{
		fmt.Println("计算结果是：",result)
	}
}

func Goroutine1(ch chan int) {
	defer close(ch)
	for index := 0; index < 100; index++ {
		time.Sleep(20*time.Microsecond)
		num:=rand.Intn(10)
		ch<-num

	}
}

func Goroutine2(ch chan int,ch2 chan  int) {
	defer close(ch2)
	for value:=range ch{
		result:=value*value*value
		ch2<-result
	}
}
