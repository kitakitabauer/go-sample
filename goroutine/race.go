package main

import (
	"fmt"
	"time"
)

func main() {
	num := 10
	go func() {
		num += 1
		fmt.Println(num)
	}()

	num += 1
	fmt.Println(num)
	time.Sleep(time.Second)
}
