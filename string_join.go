package main

import (
	"bytes"
	. "fmt"
)

func main() {
	var buf bytes.Buffer

	strs := "1"
	strs += "\n2\n"
	buf.WriteString(strs)

	Printf("\"" + buf.String() + "\"")
}
