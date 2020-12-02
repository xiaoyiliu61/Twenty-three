package main

import (
	"fmt"
	"strings"
)

func main() {
	str:="hello hello hello world"
    fmt.Println(strings.Count(str,"llo"))
}
