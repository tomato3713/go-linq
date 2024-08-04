package main

import (
	"fmt"
	"github.com/tomato3713/linq"
	"slices"
)

func main() {
	list := []int{5, 3, 10, 1, 33, 4, 12, 11, 15}

	cond1 := func(v int) bool { return v > 5 }
	cond2 := func(v int) bool { return v < 30 }

	cmp := func(a, b int) int { return a - b }

	selector := func(item int) int { return item * 2 }

	q := linq.New(slices.Values(list))

	for v := range q.
		Where(cond1).      // -> []int{10, 33, 12, 11, 15}
		Where(cond2).      // -> []int{10, 12, 11, 15}
		OrderBy(cmp).      // -> []int{10, 11, 12, 15}
		Select(selector) { // -> []int{20, 22, 24, 30}
		// output
		// 20
		// 22
		// 24
		// 30
		fmt.Println(v)
	}
}
