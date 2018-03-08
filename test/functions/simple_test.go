package functions

import (
	"testing"
	"function"
)


func TestIsPowerOfFour(t *testing.T) {
	if function.IsPowerOfFour(5) {
		t.Errorf("5 should not be power of 4.")
	}

	if !function.IsPowerOfFour(16) {
		t.Errorf("16 should be power of 4.")
	}

	if function.IsPowerOfFour(32) {
		t.Errorf("32 should not be power of 4.")
	}

	if !function.IsPowerOfFour(-65536) {
		t.Errorf("-65536 should be power of 4.")
	}
}
