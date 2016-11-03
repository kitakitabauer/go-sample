package main

import (
	. "fmt"
	"strconv"
	"bytes"
)

type Vertex struct {
	x string
	y string
}

func main() {
	var m = map[string]int{
		"Bell Labs": 1,
		"Google": 2,
	}
	Println(m) // map[Bell Labs:1 Google:2]

	Println(m["Bell Labs"]) // 1

	delete(m, "Google")

	Println(m) // map[Bell Labs:1]

	// キーの有無✔
	str, ok := m["Bell Labs"]
	Println(str, ok) // 1 true

	int := 0
	Println("fill前：", strconv.Itoa(int))
	fill(&int)
	Println("fill後：", strconv.Itoa(int))
	// reset
	int = 0

	var buf bytes.Buffer

	buf.WriteString("abc")
	Println("fillBuf前：", buf.String())
	fillBuf(&int, &buf)
	Println("fillBuf後：", buf.String())
	// reset
	int = 0

	var buf1, buf2 *bytes.Buffer
	buf1 = bytes.NewBufferString("")
	buf2 = bytes.NewBufferString("")
	var bufMap = map[string]*bytes.Buffer{"buf1":buf1, "buf2":buf2}

	fillBufMap(&int, bufMap)
	Println("fillBufMap後：", bufMap)
}

func fill(counter *int) {
	for *counter < 5 {
		*counter++
	}
}

func fillBuf(counter *int, buf *bytes.Buffer) {
	Println(buf.String())
	for *counter < 5 {
		buf.WriteString(strconv.Itoa(*counter))
		*counter++
	}
}

func fillBufMap(counter *int, bufMap map[string]*bytes.Buffer) {
	for *counter < 5 {
		bufMap["buf1"].WriteString(strconv.Itoa(*counter))
		bufMap["buf2"].WriteString(strconv.Itoa(*counter))
		*counter++
	}
}
