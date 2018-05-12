package lesson2

import (
	"testing"
)

func TestSolution2(t *testing.T) {
	A := []int{9, 3, 9, 3, 9, 7, 9}
	result := Solution2(A)
	if result != 7 {
		t.Errorf("expect 7, but found %v", result)
	}
}