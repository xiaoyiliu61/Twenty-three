package main

import (
	"fmt"
	"sync"
	"time"
)

//模拟工厂生产和消费者消费的场景。
//有一个工厂，生产日用品，每8ms生产一件商品，然后进入库存；
//另外有三个销售渠道来销售这批产品，三个销售渠道均是每20ms销售一件商品。
//当库存达到300件时，仓库满了，停止生产程序结束；
//当库存为0时，没有库存，程序也停止执行。编程模拟实现上述场景。
//4.23号作业
/*var ch chan bool
var goods int
var  wg1 sync.WaitGroup
var  mutex sync.Mutex
func main() {

	ch=make(chan bool)
	go PrintNum1()

	rs:=<-ch
	if rs==true {
		fmt.Println("停止生产...")
	}

	wg1.Add(3)
	go JD()
	go TaoBao()
    go Officer()
	wg1.Wait()
	fmt.Printf("剩余%d件商品，没有库存了，活动结束。。",goods)
}

func PrintNum1() {
	for {
		time.Sleep(8*time.Millisecond)
		fmt.Printf("生产%d件商品\n",goods)
		goods++
		if goods>=300 {
			ch<-true
		}
	}
}

func JD()  {
	for {
		mutex.Lock()
		if goods>0 {
			time.Sleep(20*time.Millisecond)
			goods--
			fmt.Printf("京东卖掉一件商品，目前还剩%d件商品\n",goods)
		}else if goods<=0 {
			wg1.Done()
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}

func TaoBao()  {
	for {
		mutex.Lock()
		if goods>0 {
			time.Sleep(20*time.Millisecond)
			goods--
			fmt.Printf("淘宝卖掉一件商品，目前还剩%d件商品\n",goods)
		}else if goods<=0 {
			wg1.Done()
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}

func Officer()  {
	for {
		mutex.Lock()
		if goods>0 {
			time.Sleep(20*time.Millisecond)
			goods--
			fmt.Printf("官网卖掉一件商品，目前还剩%d件商品\n",goods)
		}else if goods<=0 {
			wg1.Done()
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}*/

var produce int
var wg1 sync.WaitGroup
var mutex sync.Mutex
func main() {
	wg1.Add(4)
  go Produce()
	go Sale("渠道1")
	go Sale("渠道2")
	go Sale("渠道3")
	wg1.Wait()
	fmt.Println("over...")
}

func Produce()  {
	defer wg1.Done()
	for {
		mutex.Lock()
		time.Sleep(8*time.Millisecond)
		if produce<300&&produce>=0 {
			produce++
			mutex.Unlock()
		}else {
			fmt.Println("仓库已满，停止生产")
			mutex.Unlock()
			break
		}
	}
}
func Sale(name string)  {
	defer wg1.Done()
	for {
		mutex.Lock()
		time.Sleep(20*time.Millisecond)
		if produce<=0 {
			fmt.Println("仓库没有库存，停止售卖")
			mutex.Unlock()
			break
		}
		produce--
		fmt.Println(name, "销售了商品，仓库剩余",name,produce)
		mutex.Unlock()
	}
}


