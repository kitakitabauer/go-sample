package array

import (
	. "fmt"
)
func main() {

	p := []int{2, 3, 4, 5, 6}

	for i := 0; i < len(p); i++ {
		Printf("p[%d] == %d\n", i, p[i])
	}

	// array section
	ans1 := p[1:4]
	ans2 := p[1:]
	ans3 := p[:4]
	Println(ans1) // [3 4 5]
	Println(ans2) // [3 4 5 6]
	Println(ans3) // [2 3 4 5]

	make1 := make([]int, 5)
	Println(make1, len(make1)) // [0 0 0 0 0] 5
	make2 := make([]int, 0, 5)
	Println(make2, len(make2), cap(make2)) // [] 0 5
}
