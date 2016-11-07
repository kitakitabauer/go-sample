package main

import (
	"os"
	"fmt"
	"log"
	"path"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		log.Fatal("Usage: ", path.Base(args[0]), " anything")
	}

	fmt.Println("len: ", len(args), " values: ", args)
}
