package trees

import (
	"github.com/play-code/src/basedata/ele"
	"github.com/play-code/src/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBfsQueue(t *testing.T) {
	root := testdata.TreeNodeData()
	got := BfsQueue(root)

	except := []ele.TreeEle{2, 1, 3, 9, 5, 4, 8}
	assert.Equal(t, except, got)
}

func TestDfs(t *testing.T) {
	root := testdata.TreeNodeData()
	got := Dfs(root)

	except := []ele.TreeEle{2, 1, 9, 8, 3, 5, 4}
	assert.Equal(t, except, got)
}
