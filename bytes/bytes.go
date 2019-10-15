package bytes

import (
	"bytes"
	. "fmt"
)

func main() {
	var buf bytes.Buffer
	Println(buf.Len())

	buf.WriteString("hogehoge")

	buf.Reset()
	Println(buf.Len())
}
