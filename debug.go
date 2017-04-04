package main

import (
	debug "github.com/tj/go-debug"
	"os"
	"fmt"
)

func init() {
	// It was not able to print...
	os.Setenv("DEBUG", "*")
}

func main() {
	var d = debug.Debug("debug test")

	fmt.Println(os.Getenv("DEBUG"))

	d("test")
}
