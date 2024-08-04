package main

import (
	"fmt"
	"github.com/tomato3713/linq"
	"slices"
	"strings"
)

func main() {
	list := []string{"Bob", "Kevin", "Stuart", "Dave", "Tim", "Carl", "Tom", "Jerry"}

	cond := func(v string) bool { return len(v) > 4 }

	q := linq.New(slices.Values(list))
	// filtering and sorted, and make a slice
	out := slices.Collect(q.Where(cond).OrderBy(strings.Compare).ToIter())

	fmt.Println(strings.Join(out, ", "))
	// output:
	// Jerry, Kevin, Stuart
}
