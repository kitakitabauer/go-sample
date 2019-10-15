package _struct

import (
	. "fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	X float64
	Y float64
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// define my type
type MyFloat float64

func(f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// initialize, and access
	v := Vertex{1, 2}
	Println(v.X)
	x, ok := v["X"].(int)
	Println(x, ok) // 1

	// pointer
	p := Vertex{1, 2}
	q := &p
	Println(q.X) // 1

	// initialize only X filed
	Println(Vertex{X: 1}) // {1 0}

	// use new initialize
	// fileds are implicitly initialized
	var newV *Vertex = new(Vertex)
	// or newV = new(Vertex)
	Println(newV) // &{0 0}

	v2 := &Vertex2{3, 4}
	Println(v2.Abs()) // 5

	f := MyFloat(-math.Sqrt(2))
	Println(f.Abs()) // 1.41421356…
}
