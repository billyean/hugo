package util

import (
	"testing"
	"util"
)

func TestAbs(t *testing.T) {
	nameWithPrefix := util.GetPrefix()

	if _, exists := nameWithPrefix["JohnDoe"]; exists {
		nameWithPrefix["JohnDoe"] = "Mr"
	} else {
		nameWithPrefix["JohnDoe"] = "Unknown"
	}

	if nameWithPrefix["JohnDoe"] != "Mr" {
		t.Errorf("John Doe is not in the orginal map.")
	}
}
