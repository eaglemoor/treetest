package treetest_test

import (
	"testing"

	. "github.com/eaglemoor/treetest"

	"github.com/stretchr/testify/assert"
)

func Test_Tree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		in   [][2]int
		out  string
	}{
		{
			name: "[1 2], [1 3], [2 4], [3 5], [7 8]",
			n:    10,
			in:   [][2]int{{1, 2}, {1, 3}, {2, 4}, {3, 5}, {7, 8}},
			out:  "sqrt(5) + sqrt(2) + 3",
		},
	}

	for _, data := range tests {
		t.Run(data.name, func(tt *testing.T) {
			out := Tree(data.n, data.in)
			assert.Equal(tt, out, data.out)
		})
	}
}
