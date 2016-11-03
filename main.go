package main

import (
	"fmt"
	"time"
	"os"
	"log"
)

var tracer = log.New(os.Stdout, "", log.Lshortfile)

func main() {
	fmt.Println("test!")
}

func Sum(x, y int) int {
	return x + y
}

func TimeZone() {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	n1 := time.Now()

	n2 := n1.UTC()
	n3 := n2.In(jst)
	now := time.Date(
		n3.Year(),
		n3.Month(),
		n3.Day(),
		0,
		0,
		0,
		0,
		jst,
	)

	nowUnix := now.Unix()

	tracer.Println("jst:", jst)
	tracer.Println("n1:", n1) // JST(日本標準時)のためジャストな時間で表示
	tracer.Println("n2:", n2) // UTC(協定世界時)で表示のため、-9時間
	tracer.Println("n3:", n3) // Asia/Tokyoに指定した時間、JSTと一緒になる
	tracer.Println("now:", now) // 実行日の0時
	tracer.Println("nowUnix:", nowUnix) // 1970/01/01 UTC からの経過時間ミリ秒

	prevDay := now.AddDate(0, 0, -1)

	tracer.Println("nowPrevDay:", prevDay)
	tracer.Println("nowPrevDayUnix:", prevDay.Unix()) // 前日の経過時間ミリ秒
}
