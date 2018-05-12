package lesson4

import (
	"testing"
)

func TestSolution3(t *testing.T) {
	A := []int{4, 1, 3, 2}
	result := Solution3(A)
	if result != 1 {
		t.Errorf("expect 1, but found %v", result)
	}

	A = []int{4, 1, 3}
	result = Solution3(A)
	if result != 0 {
		t.Errorf("expect 0, but found %v", result)
	}

	A = []int{1, 4, 1}
	result = Solution3(A)
	if result != 0 {
		t.Errorf("expect 0, but found %v", result)
	}
}