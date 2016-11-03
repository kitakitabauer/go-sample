package main

import (
	"time"
	. "fmt"
)

const (
	monthToMillisec = 2592000
)

func main() {
	n := time.Now() // 現在時刻
	Println(n)
	Println(n.Day())
	Println(n.ISOWeek())
	Println(n.Unix())
	Println(n.Location())
	Println(n.MarshalJSON())

	dateFormat := "2006-01-02"
	targetDayStr := "2016-10-19"

	parseDate, err := time.Parse(dateFormat, targetDayStr)
	if err != nil {
		panic(err)
	}

	Println("parseDate:", parseDate) // 2016-10-19 00:00:00 +0000 UTC

	nowDate, err := time.Parse(dateFormat, time.Now().Format(dateFormat))
	if err != nil {
		panic(err)
	}

	Println("nowDate:", nowDate) // 2016-10-19 00:00:00 +0000 UTC

	// judge future
	Println(nowDate.After(parseDate))

	location, err := time.LoadLocation("Japan")
	if err != nil {
		panic(err)
	}

	var from, to time.Time
	from = time.Date(2016, 9, 1, 0, 0, 0, 0, time.UTC)
	to = time.Date(2016, 9, 1, 1, 0, 0, 0, time.UTC)

	fromLoc := time.Unix(from.Unix(), 0).In(location)
	toLoc := time.Unix(to.Unix(), 0).In(location)
	Println("fromLoc", fromLoc)
	Println("toLoc", toLoc)
	Println("fromLoc.Unix()", fromLoc.Unix()) // 1472688000
	Println("toLoc.Unix()", toLoc.Unix()) // 1475280000

	Println("diffLoc", toLoc.Unix() - fromLoc.Unix()) // 2592000
	//fromTime := time.Unix(from, 0).In(location)
	//toTime := time.Unix(to, 0).In(location)
	diffMonth := diffMonth(toLoc, fromLoc)
	Println("diffMonth:", diffMonth)
}

// In consideration of exceeding years,
// calculate the months from created date to expired date.
func diffMonth(expiresTime, createdTime time.Time) int {
	return (int(expiresTime.Unix()) - int(createdTime.Unix())) / monthToMillisec
}

// In consideration of exceeding years,
// calculate the months from created date to expired date.
func diffMonthBad(expiresTime, createdTime time.Time) int {
	diffYear := expiresTime.Year() - createdTime.Year()
	return 12*diffYear + int(expiresTime.Month()) - int(createdTime.Month())
}
