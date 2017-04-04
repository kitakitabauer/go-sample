package main

import "fmt"

func main() {

	// for zero fill
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("%04d", i))
	}

}

