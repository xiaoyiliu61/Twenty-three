package main

import (
	"fmt"
	"strings"
)

func main() {
    str:="zhongguojiayouwuhanjiayou"
    stringChu(str)

}
func stringChu(n string) string {
	str1:=strings.Split(n,"")
	map2:=make(map[string]int)
	for index := 0; index < len(str1); index++ {
		map2[str1[index]]++
	}
	fmt.Println(map2)
    return n
}
