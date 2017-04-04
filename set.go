package main

import (
	. "fmt"
)

var (
	set = map[string]struct{}{}
)

func main() {
	set["1"] = struct{}{}

	_, ok := set["1"]
	Println(ok)
	_, ok = set["2"]
	Println(ok)
}
