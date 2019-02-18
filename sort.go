package main

import (
	"fmt"
	"sort"
)

type Food struct {
	Name  string
	Price int
}

func main() {
	foods := make([]Food, 4)
	foods[0] = Food{Name: "りんご", Price: 120}
	foods[1] = Food{Name: "みかん", Price: 150}
	foods[2] = Food{Name: "バナナ", Price: 100}
	foods[3] = Food{Name: "パイナップル", Price: 100}

	sort.Slice(foods, func(i, j int) bool {
		return foods[i].Price > foods[j].Price
	})

	fmt.Printf("%+v", foods)
}
