package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("hoge%s\n", "hoge")
	fmt.Printf("hoge2%s\n", "hoge2")
	log.Printf("fuga%s\n", "fuga")
	log.Printf("fuga2%s\n", "fuga2")
	fmt.Println("end")
}
