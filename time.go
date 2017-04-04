package main

import (
	"time"
	. "fmt"
)

func main() {
	start := time.Now();

	time.Sleep(1000 * time.Millisecond)

	end := time.Now();
	Printf("%f秒\n",(end.Sub(start)).Seconds())


	//startAt := int64(14000000)
	startAt := time.Now().Unix()
	expires := time.Unix(startAt, int64(0)).Add(time.Duration(60)*time.Second)

	Printf("%v", expires)
}
