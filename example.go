package treetest

import (
	"fmt"
	"strings"
)

func Tree(n int, data [][2]int) string {
	tree := make([]*int, n)
	first := make(map[int]bool, n)

	for _, row := range data {
		if tree[row[0]] == nil {
			k := 1
			tree[row[0]] = &k
			first[row[0]] = true
		}

		if tree[row[1]] == nil {
			tree[row[1]] = tree[row[0]]
		}

		*tree[row[1]]++
	}

	s := make([]string, 0, n)
	all1 := 0
	for k, i := range tree {
		if i == nil {
			all1++
			continue
		}
		if isFirst := first[k]; !isFirst {
			continue
		}

		s = append(s, fmt.Sprintf("sqrt(%d)", *i))
	}
	if all1 > 0 {
		s = append(s, fmt.Sprintf("%d", all1))
	}

	return strings.Join(s, " + ")
}
