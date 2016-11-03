package main
import (
	. "fmt"
)

func main() {
	var slice1 []int
	slice2 := []string{"a", "b", "c"}
	Println(slice1) // []
	Println(slice2) // [a b c]

	// slice to new slice
	a := [5]int{1, 2, 3, 4, 5}
	b := a[2:4] // a and b is the same memory space
	Println(a) // [1 2 3 4 5]
	Println(b) // [3 4]

	// array to new slice
	array := []string{"a", "b", "c"}
	slice3 := array[:]
	Println(array) // [a b c]
	Println(slice3) // [a b c]

	// use make
	var slice4 []int = make([]int, 3, 3) // type, length, size
	Println(slice4) // [0 0 0]
	Println(len(slice4), cap(slice4)) // 3 3

	// assign by loop
	s := make([]byte, 5)
	t := make([]byte, len(s), (cap(s)+1)*2)
	for i := range s {
		t[i] = s[i]
	}
	Println(s, t) // [0 0 0 0 0] [0 0 0 0 0]

	// use copy function
	u := make([]byte, 5)
	v := make([]byte, len(s), (cap(s)+1)*2) // larger of cap
	Println(copy(u, v)) // 5
	Println(u, v) // [0 0 0 0 0] [0 0 0 0 0]

	// use append
	// expand automatic
	var slice5 []int = make([]int, 3)
	Println(slice5) // [0 0 0]
	slice5 = append(slice5, 10, 20, 30, 40, 50)
	Println(slice5) // [0 0 0 10 20 30 40 50]
}
