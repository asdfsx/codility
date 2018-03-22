package lesson4

import (
	"testing"
	"fmt"
)

func TestSolution4(t *testing.T) {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	result := Solution4(5, A)
	if !sliceCompare(result, []int{3, 2, 2, 4, 2}) {
		t.Errorf("expect [3, 2, 2, 4, 2], but found %s", result)
	}
}

func sliceCompare(A []int, B []int) bool {
	if len(A) != len(B){
		fmt.Println(len(A), len(B))
		return false
	}

	for i, v := range A{
		if v != B[i]{
			return false
		}
	}
	return true
}