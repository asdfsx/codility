package lesson4

import (
	"testing"
)

func TestSolution(t *testing.T) {
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	result := Solution(5, A)
	if result != 6{
		t.Errorf("expect 6, but found %s", result)
	}

	A = []int{2, 2, 2, 2, 2}
	result = Solution(2, A)
	if result != -1{
		t.Errorf("expect -1, but found %s", result)
	}
}
