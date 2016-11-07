package main

import (
	"os"
	"log"
	"path"
	"strconv"
	"fmt"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: ", path.Base(args[0]), " number")
	}

	if val, err := strconv.Atoi(args[1]); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("val: ", val)
	}
}
