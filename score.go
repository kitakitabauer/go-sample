package main

import (
	"fmt"
	"math"
)

const (
	assignTotal     = 1000

	highPopularity = 100
	lowPopularity  = 1
)

// ニューリリース含め同一クラスタで処理する場合
func main() {
	popularity := highPopularity
	assignNum := getAssignOfNum()
	fmt.Println("assignNum: ", assignNum)

	for i := 1; i <= assignTotal; i++ {
		if popularity < lowPopularity {
			fmt.Printf("Unexpected:  %d\n", i)
			continue
		}

		if i%assignNum == 0 {
			fmt.Println("popularity: ", popularity)
			popularity--
		}
	}
}

func getAssignOfNum() int {
	f := math.Ceil(float64(assignTotal) / (highPopularity - lowPopularity + 1))
	fmt.Println("assignNum(float): ", f)
	return int(f)
}
