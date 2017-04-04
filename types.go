package main

import (
	"fmt"
	"go/types"
)

type Hoge struct {
	fuga string
}

func main() {
	hoge := Hoge{fuga: "fuga"}

	fmt.Printf("%v", types.Type(hoge))

}
