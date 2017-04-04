package main

import (
	"strings"
	"fmt"
)

func main() {
	var test = "hogehoge\""
	ret := strings.TrimRight(test, "\"")
	fmt.Println(ret)
}
