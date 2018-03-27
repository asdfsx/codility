package lesson6

import (
	"testing"
)

func TestSolution2(t *testing.T) {
	A := []int{2, 1, 1, 2, 3, 1}
	result := Solution2(A)
	if result != 3 {
		t.Errorf("expect 3, but found %s", result)
	}

	A = []int{0}
	result = Solution2(A)
	if result != 1 {
		t.Errorf("expect 1, but found %s", result)
	}
}