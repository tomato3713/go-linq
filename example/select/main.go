package main

import (
	"github.com/tomato3713/linq"
	"slices"
	"fmt"
)

func main() {
	list := []int{5, 3, 10, 1, 33, 4, 12, 11, 15}

	cond1 := func(v int) bool { return v > 5 }
	cond2 := func(v int) bool { return v < 30 }

	cmp := func(a, b int) int { return a-b }

	selector := func(item int) int { return item*2 }

	q := linq.New(slices.Values(list))

	for v := range q.Where(cond1).Where(cond2).OrderBy(cmp).Select(selector) {
		fmt.Println(v)
	}
}
