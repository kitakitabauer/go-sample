package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	files, _ := filepath.Glob("./*.go")
	for _, v := range files {
		fmt.Println(v)
	}
}
