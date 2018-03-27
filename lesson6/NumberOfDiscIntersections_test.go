package lesson6

import (
	"testing"
)

func TestSolution4(t *testing.T) {
	A := []int{1, 5, 2, 1, 4, 0}
	result := Solution4(A)
	if result != 11 {
		t.Errorf("expect 11, but found %s", result)
	}
}