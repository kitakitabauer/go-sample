package main

import (
	"bytes"
	. "fmt"
	"strconv"
)

type Vertex struct {
	x string
	y string
}

func main() {
	var m = map[string]int{
		"Bell Labs": 1,
		"Google":    2,
		"":          3,
	}
	Println(m)              // map[Bell Labs:1 Google:2 :3]
	Println(m["Bell Labs"]) // 1
	Println(m[""])          // 3

	Println(m) // map[Bell Labs:1 :3]

	// キーの有無✔
	num, ok := m["Bell Labs"]
	Println(num, ok) // 1 true

	num, ok = m["hoge"]
	Println(num, ok) // 0 false

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
	var bufMap = map[string]*bytes.Buffer{"buf1": buf1, "buf2": buf2}

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
