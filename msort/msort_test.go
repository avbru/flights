package msort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSort(t *testing.T) {
	n := 10000
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()
	}

	require.IsIncreasing(t, MergeSort(arr))
}
