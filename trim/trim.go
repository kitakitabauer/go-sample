package trim

import (
	"fmt"
	"strings"
)

func main() {
	var test = "hogehoge\""
	ret := strings.TrimRight(test, "\"")
	fmt.Println(ret)

	fmt.Println(strings.Trim("\"already.used.code", "\"\\"))
}
