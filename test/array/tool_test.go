package array

import (
	"testing"
	"array"
)

func TestCount(t *testing.T) {
	bits := []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2}

	if len(array.Count(17)) != len(bits) {
		t.Errorf("Output len is not same as expected.")
	}
}
