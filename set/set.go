package set

import (
	. "fmt"

	"github.com/deckarep/golang-set"
)

var (
	set = map[string]struct{}{}
)

func main() {
	//native()

	makeSetByGolangSet()
}

func native() {
	set["1"] = struct{}{}

	_, ok := set["1"]
	Println(ok)
	_, ok = set["2"]
	Println(ok)
}

func makeSetByGolangSet() {
	s1 := mapset.NewSet()
	s1.Add("1")
	s1.Add("2")
	s1.Add(3)
	s1.Add(4)
	print(s1) // Set{4, "1", "2", 3}

	slice := []interface{}{"1", "2", 3, "5"}
	s2 := mapset.NewSetFromSlice(slice)
	print(s2) // Set{"5", "1", "2", 3}

	all := s1.Union(s2)
	print(all)                // Set{3, 4, "1", "2", "5"}
	print(all.Difference(s1)) // Set{"5"}

	// size
	Println(s1.Cardinality())  // 4
	Println(s2.Cardinality())  // 4
	Println(all.Cardinality()) // 5

	// check
	Println(s1.Contains(1))   // false
	Println(s1.Contains("1")) // true

	// remove
	s1.Remove(4)
	print(s1) // Set{"1", "2", 3}

	// clear
	s1.Clear()
	print(s1) // Set{}
}

func print(s mapset.Set) {
	Println(s.String())
}
