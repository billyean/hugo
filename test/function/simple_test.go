package function

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

func TestAbs(t *testing.T) {
	if function.Abs(-1) != 1 {
		t.Errorf("Abs(-1) is expected to 1 but actual is %d.", function.Abs(-1))
	}

	if function.Abs(0) != 0 {
		t.Errorf("Abs(0) is expected to 0 but actual is %d.", function.Abs(0))
	}

	if function.Abs(1) != 1 {
		t.Errorf("Abs(1) is expected to 1 but actual is %d.", function.Abs(1))
	}
}
