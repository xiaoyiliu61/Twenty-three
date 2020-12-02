package main

import "fmt"

func main() {
	var height int
	var weight int
	fmt.Printf("身高和体重是：")
	fmt.Scanln(&height,&weight)
    fmt.Printf("身高是：%dcm\n体重是:%dkg",height,weight)
}
