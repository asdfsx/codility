package lesson7

import (
	"testing"
)

func TestSolution3(t *testing.T) {
	A := []int{4, 3, 2, 1, 5}
	B := []int{0, 1, 0, 0, 0}
	result := Solution3(A, B)
	if result != 2 {
		t.Errorf("expect 2, but found %s", result)
	}
}