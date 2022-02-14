package sorting

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSorts(t *testing.T) {
	input := [][]int{
		{1, 22, 3, 15, 9, 8},
	}

	except := [][]int{
		{1, 3, 8, 9, 15, 22},
	}

	testFunc := []func([]int){
		insert, selectSort, shell,
	}

	for i, sort := range testFunc {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for i := range input {
				temp := make([]int, len(input[i]))
				copy(temp, input[i])
				sort(temp)
				assert.Equal(t, except[i], temp)
			}
		})
	}
}
