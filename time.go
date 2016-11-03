package main

import (
	"time"
	. "fmt"
)

func main() {
	start := time.Now();

	time.Sleep(5000 * time.Millisecond)

	end := time.Now();
	Printf("%f秒\n",(end.Sub(start)).Seconds())
}
