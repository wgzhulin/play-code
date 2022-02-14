package testdata

import (
	"github.com/play-code/src/basic/datastructure/basedata/ele"
	"github.com/play-code/src/basic/datastructure/trees"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBinaryTreeData(t *testing.T) {
	input := []ele.TreeEle{2, 1, 3, 9, 0, 5, 4, 8}
	root := NewBinaryTreeData(input)

	assert.Equal(t, clearZero(input), trees.BfsQueue(root))
}

func clearZero(s []ele.TreeEle) []ele.TreeEle {
	result := make([]ele.TreeEle, 0, len(s))
	for i := range s {
		if s[i] == 0 {
			continue
		}
		result = append(result, s[i])
	}

	return result
}
