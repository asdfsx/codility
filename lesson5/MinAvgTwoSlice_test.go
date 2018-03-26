package lesson5

import (
	"testing"
)

func TestSolution4(t *testing.T) {
	A := []int{4, 2, 2, 5, 1, 5, 8}
	result := Solution4(A)
	if result != 1{
		t.Errorf("expect 1, but found %s", result)
	}
}