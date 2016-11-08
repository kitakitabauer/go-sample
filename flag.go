package main

import (
	"flag"
	"os"
	. "fmt"
	"github.com/kitakitabauer/go-sample/flagout"
)

var start int

func main() {
	Println(os.Args)

	flag.IntVar(&start, "start", 0, "start day")
	end := flag.Int("end", 0, "end day")
	test := flag.String("test", "", "test2")

	flagout.Flag()

	// 全てのflagを定義する前にParseするとnot definedエラーとなる
	// Parseは何度呼んでもpanicにはならない
	flag.Parse()

	Println(start)
	Println(*end)
	Println(*test)
	Println(flagout.Test2)
}
