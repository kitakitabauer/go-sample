package main

import (
	"fmt"
	"time"
)

const (
	DateFormat = "2006-01-02"
	Location   = "Asia/Tokyo"
	Timezone   = 9 * 60 * 60
)

var (
	date201706010000 = time.Date(2017, time.June, 01, 0, 0, 0, 0, time.FixedZone(Location, Timezone))
	date201707010000 = time.Date(2017, time.July, 01, 0, 0, 0, 0, time.FixedZone(Location, Timezone))
)

func main() {
	fmt.Println("20170601:", date201706010000.Unix())
	fmt.Println("20170701:", date201707010000.Unix())
}
