package adddate

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	period0531 := time.Date(2017, time.May, 31, 12, 0, 0, 0, loc)
	fmt.Println(period0531)
	fmt.Println("1ヶ月後:", period0531.AddDate(0, 1, 0))
	fmt.Println("2ヶ月後:", period0531.AddDate(0, 2, 0))
	fmt.Println("3ヶ月後:", period0531.AddDate(0, 3, 0))
	fmt.Println("4ヶ月後:", period0531.AddDate(0, 4, 0))
}
