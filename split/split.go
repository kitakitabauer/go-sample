package split

import (
	"fmt"
	"strings"
)

func main() {
	e1 := "api.request: 403::hoge::fuga"
	sp := strings.Split(e1, "api.request: ")
	fmt.Printf("%#v", sp)
	var codeStr string
	if len(sp) > 1 {
		codeStr = strings.Split(sp[1], "::")
	}
	fmt.Printf("%#v", codeStr)

	e2 := "api.request: 403"
	sp = strings.SplitAfter(e2, "api.request: ")
	fmt.Printf("%#v", sp)
}
