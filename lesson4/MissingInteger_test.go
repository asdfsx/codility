package lesson4

import (
	"testing"
)

func TestSolution2(t *testing.T) {
	A := []int{1, 3, 6, 4, 1, 2}
	result := Solution2(A)
	if result != 5{
		t.Errorf("expect 5, but found %v", result)
	}

	A = []int{1, 2, 3}
	result = Solution2(A)
	if result != 4{
		t.Errorf("expect 4, but found %v", result)
	}

	A = []int{-1, -3}
	result = Solution2(A)
	if result != 1{
		t.Errorf("expect 1, but found %v", result)
	}
}