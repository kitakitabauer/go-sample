package now

import (
	. "fmt"
	"time"

	"github.com/jinzhu/now"
)

func main() {
	Println(time.Now())
	Println()

	prints(
		now.BeginningOfYear(),
		now.BeginningOfQuarter(),
		now.BeginningOfMonth(),
		now.BeginningOfWeek(),
		now.BeginningOfDay(),
		now.BeginningOfHour(),
		now.BeginningOfMinute(),
	)

	Println()

	prints(
		now.EndOfYear(),
		now.EndOfQuarter(),
		now.EndOfMonth(),
		now.EndOfWeek(),
		now.EndOfDay(),
		now.EndOfHour(),
		now.EndOfMinute(),
	)

	Println()

	loc, _ := time.LoadLocation("Asia/Tokyo")
	t := now.New(time.Date(2017, time.May, 31, 15, 0, 0, 0, loc))

	prints(
		t.BeginningOfYear(),
		t.BeginningOfQuarter(),
		t.BeginningOfMonth(),
		t.BeginningOfWeek(),
		t.BeginningOfDay(),
		t.BeginningOfHour(),
		t.BeginningOfMinute(),
	)

	Println()

	prints(
		t.EndOfYear(),
		t.EndOfQuarter(),
		t.EndOfMonth(),
		t.EndOfWeek(),
		t.EndOfDay(),
		t.EndOfHour(),
		t.EndOfMinute(),
	)
}

func prints(times ...time.Time) {
	for i := range times {
		Println(times[i])
	}
}
