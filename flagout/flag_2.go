package flagout

import (
	"flag"
	. "fmt"
)

var (
	Test2 string
)

func Flag() {
	Println("flag_2 init!")

	flag.StringVar(&Test2, "test2", "", "test2 string")

	//flag.Parse()
	Println(Test2)
}
