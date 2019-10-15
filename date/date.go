package date

import (
	. "fmt"
	"time"
)

const (
	monthToMillisec = 2592000
	location        = "Asia/Tokyo"
	timezone        = 9 * 60 * 60
	dateFormat      = "2006-01-02"
	dateTimeFormat  = "2006-01-02 15:04:05"
	targetDayStr    = "2016-10-19"
)

func main() {
	n := time.Now() // 現在時刻
	Println(n)
	Println(n.Day())
	Println(n.ISOWeek())
	Println(n.Unix())
	Println(n.Location())
	Println(n.MarshalJSON())

	checkAddDate()

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
	Println("toLoc.Unix()", toLoc.Unix())     // 1475280000

	Println("diffLoc", toLoc.Unix()-fromLoc.Unix()) // 2592000
	//fromTime := time.Unix(from, 0).In(location)
	//toTime := time.Unix(to, 0).In(location)
	diffMonth := diffMonth(toLoc, fromLoc)
	Println("diffMonth:", diffMonth)

	newDate()
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

func newDate() time.Time {
	zone := time.FixedZone(location, timezone)
	day20160101 := time.Date(2016, 1, 1, 0, 0, 0, 0, zone)
	Println(day20160101.Unix())
	return day20160101
}

func checkAddDate() {
	date := time.Date(2017, time.March, 28, 12, 0, 0, 0, time.FixedZone(location, timezone))
	Println(date.Format(dateTimeFormat))                    // 2017-03-28 12:00:00
	Println(date.AddDate(0, -1, 0).Format(dateTimeFormat))  // 2017-02-28 12:00:00
	Println(date.AddDate(0, 0, -30).Format(dateTimeFormat)) // 2017-02-26 12:00:00

	date = time.Date(2017, time.March, 29, 12, 0, 0, 0, time.FixedZone(location, timezone))
	Println(date.Format(dateTimeFormat))                    // 2017-03-29 12:00:00
	Println(date.AddDate(0, -1, 0).Format(dateTimeFormat))  // 2017-03-01 12:00:00
	Println(date.AddDate(0, 0, -30).Format(dateTimeFormat)) // 2017-02-27 12:00:00

	date = time.Date(2017, time.May, 31, 12, 0, 0, 0, time.FixedZone(location, timezone))
	Println(date.Format(dateTimeFormat))                    // 2017-05-31 12:00:00
	Println(date.AddDate(0, -1, 0).Format(dateTimeFormat))  // 2017-05-01 12:00:00
	Println(date.AddDate(0, 0, -30).Format(dateTimeFormat)) // 2017-05-01 12:00:00

	date = time.Date(2017, time.July, 1, 12, 0, 0, 0, time.FixedZone(location, timezone))
	Println(date.Format(dateTimeFormat))                    // 2017-07-01 12:00:00
	Println(date.AddDate(0, -1, 0).Format(dateTimeFormat))  // 2017-06-01 12:00:00
	Println(date.AddDate(0, 0, -30).Format(dateTimeFormat)) // 2017-06-01 12:00:00
	Println(date.AddDate(0, 0, -31).Format(dateTimeFormat)) // 2017-05-31 12:00:00
	Println(date.AddDate(0, 0, -32).Format(dateTimeFormat)) // 2017-05-30 12:00:00

	date = time.Date(2017, time.January, 0, 23, 59, 59, 999999999, time.FixedZone(location, timezone))
	Println(date.Format(dateTimeFormat)) // 2016-12-31 23:59:59
}
