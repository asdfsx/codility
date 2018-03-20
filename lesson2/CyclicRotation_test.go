package lesson2

import (
	"testing"
)

func TestSolution(t *testing.T) {
	A := []int{3, 8, 9, 7, 6}
	K := 3
	result := Solution(A, K)
	if !compare(result, []int{9, 7, 6, 3, 8}) {
		t.Errorf("expect [9, 7, 6, 3, 8], but found %s", result)
	}

	A = []int{0, 0, 0}
	K = 1
	result = Solution(A, K)
	if !compare(result, []int{0, 0, 0}) {
		t.Errorf("expect [0, 0, 0], but found %s", result)
	}

	A = []int{1, 2, 3, 4}
	K = 4
	result = Solution(A, K)
	if !compare(result, []int{1, 2, 3, 4}) {
		t.Errorf("expect [1, 2, 3, 4], but found %s", result)
	}

	A = []int{}
	K = 1
	result = Solution(A, K)
	if !compare(result, nil) {
		t.Errorf("expect nil, but found %s", result)
	}
}

func compare(A []int, B[] int) bool{
	if A == nil && B == nil{
		return true
	}
	if len(A) != len(B){
		return false
	}
	for i, v := range A{
		if v != B[i]{
			return false
		}
	}
	return true
}