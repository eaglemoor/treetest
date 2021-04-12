package treetest

import (
	"fmt"
	"strings"
)

func Tree(n int, data [][2]int) string {
	tree := make(map[int]*int, len(data))
	first := make([]int, 0, len(data))

	for _, row := range data {
		if _, ok := tree[row[0]]; !ok {
			k := 1
			tree[row[0]] = &k
			first = append(first, row[0])
		}

		if _, ok := tree[row[1]]; !ok {
			k := tree[row[0]]
			tree[row[1]] = k
		}

		*tree[row[1]]++
	}

	for i := 1; i <= n; i++ {
		if _, ok := tree[i]; !ok {
			k := 1
			tree[i] = &k
			first = append(first, i)
		}
	}

	fmt.Println(tree)
	fmt.Printf("%+v\n", first)

	s := make([]string, 0, len(first))
	all1 := 0
	for _, i := range first {
		val := *tree[i]
		if val == 1 {
			all1++
			continue
		}

		s = append(s, fmt.Sprintf("sqrt(%d)", val))
	}
	if all1 > 0 {
		s = append(s, fmt.Sprintf("%d", all1))
	}

	return strings.Join(s, " + ")
}
