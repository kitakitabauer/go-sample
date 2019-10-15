package _continue

import (
	"fmt"
	"strconv"
)

var m = map[int]struct{}{
	3: struct{}{},
}

func exclude(j int) bool {
	_, found := m[j]
	fmt.Println(found)
	return found
}

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if exclude(j) {
				fmt.Println("j = " + strconv.Itoa(j))
				continue
			}
		}
	}
}
