package common

import (
	"testing"
)

func TestRange(t *testing.T) {
	testrange := Range(1, 5)
	expected := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(testrange); i++ {
		if testrange[i] != expected[i] {
			t.Error("Created test range is incorrect")
		}
	}
}
