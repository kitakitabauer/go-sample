package main

import (
	"strconv"
	"fmt"
)

func main() {
	ids := create()

	var once = 100
	var counter = 0


	for counter < len(ids) {
		fmt.Println("before")
		fmt.Printf("counter: %d\n", counter)

		work := ids[counter:counter+once]
		counter += once

		fmt.Println("after")
		fmt.Printf("counter: %d\n", counter)
		fmt.Println(work)
	}
}

// 1000までのidListを作る
func create() (ids []string) {
	ids = make([]string, 1000)

	for i := 0; i < 1000; i++ {
		ids[i] = strconv.Itoa(i)
	}

	return

}
