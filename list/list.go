package list

import (
	"container/list"
	"fmt"
)

func stacktest() {
	s := list.New()
	fmt.Print("--Stack--\nPush: ")
	for i := 0; i < 10; i++ {
		s.PushBack(i)
		fmt.Print(i, ",")
	}
	fmt.Print("\nLen: ", s.Len())
	fmt.Print("\nPop : ")
	for i := 0; i < 10; i++ {
		r := s.Remove(s.Front())
		fmt.Print(r, ",")
	}
	fmt.Print("\nLen: ", s.Len())
}

func test() {
	var idx []int

	fmt.Print(len(idx))
}

func main() {
	stacktest()
	test()
}
