package lesson6

import (
	"testing"
)

func TestSolution3(t *testing.T) {
	A := []int{-3, 1, 2, -2, 5, 6}
	result := Solution3(A)
	if result != 60 {
		t.Errorf("expect 60, but found %s", result)
	}

	A = []int{-10, 2, 4}
	result = Solution3(A)
	if result != -80 {
		t.Errorf("expect -80, but found %s", result)
	}
}