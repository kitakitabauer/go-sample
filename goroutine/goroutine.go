package main

import (
	"log"
	"os"
	"time"
)

func logging(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}

func main() {
	name := "Gopher"

	go logging("Hello, %s\n", name)

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
	go logger.Printf("Hello, %s\n", name)

	go func() {
		log.Printf("Hello, %s\n", name)
	}()

	time.Sleep(1000 * time.Millisecond)
}
