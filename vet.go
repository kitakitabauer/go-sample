package main

import (
	"fmt"
)

func main() {
	return
	fmt.Println("Hello, World!") // unreachable code
}

// % go vet vet.go                                                                                                                            20:51:30 [0]
// vet.go:9: unreachable code
// exit status 1
