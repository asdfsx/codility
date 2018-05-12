package lesson3

import (
	"testing"
)

func TestSolution(t *testing.T) {
 	A:= []int{3,1,2,4,3}
	result := Solution(A)
	if result != 1{
		t.Errorf("expect 1, but found %v", result)
	}

	A = []int{1000,1000}
	result = Solution(A)
	if result != 0{
		t.Errorf("expect 0, but found %v", result)
	}

	A = []int{1,-1,1,-1}
	result = Solution(A)
	if result != 0{
		t.Errorf("expect 0, but found %v", result)
	}
}