package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
//在main中启动3个子协程，执行一些耗时操作。
//使用sync.WaitGroup等待子协程任务完全执行结束后，整个程序执行结束
func main() {
	wg.Add(3)
    go one()

    go two()

	go three()
	wg.Wait()
	fmt.Println("程序执行结束")
}
func one()  {
	arr:=[2]string{"L","Y"}
	for index := 0; index < 999; index++ {
		fmt.Printf("字符%s\n",arr[index%2])
		time.Sleep(99*time.Microsecond)
	}
	wg.Done()
}

func two() {
	for index := 0; index < 999; index++ {
		fmt.Printf("数字%d\n",index)
		time.Sleep(99*time.Millisecond)
	}
	wg.Done()
}

func three() {
   arr1:=[3]string{"I","love","you"}
	for index := 0; index < 999; index++ {
		fmt.Printf("%s\n",arr1[index%3])
		time.Sleep(99*time.Microsecond)
	}
	wg.Done()
}
