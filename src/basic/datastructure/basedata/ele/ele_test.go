package ele

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEleEquals(t *testing.T) {
	ele := NewEle(1)
	ele1 := NewEle(1)

	ele2 := NewEle(2)
	assert.True(t, ele.Equals(ele1))
	assert.False(t, ele.Equals(ele2))
}

