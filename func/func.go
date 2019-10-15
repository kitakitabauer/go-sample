package _func

import . "fmt"

func main() {
	Println(f())

	// can be assigned to variable
	fn := func(x, y float64) (float64, float64) {
		return x, y
	}

	var a, b float64
	x, y := fn(a, b)

	Println(x, y)

	Println(adder())

	Println(retVal())
}

func f() (s string) {
	s = "aaa"

	return // don't need return
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func retVal() (s1, s2, s3 string) {
	return "a", "b", "c"
}
