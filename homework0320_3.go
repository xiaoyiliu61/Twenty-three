package main

import (
	"fmt"
	"strings"
)

func main() {
   str:="WUHANJIAYOUZHONGGUOJIAYOU"
   b:=strings.Contains(str,"ZHONGGUO")
   fmt.Println(b)
   if b==true{
	   index:=strings.Index(str,"ZHONGGUO")
	   fmt.Println(index)
   }

}
