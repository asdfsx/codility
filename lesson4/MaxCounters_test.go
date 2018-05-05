package lesson4

import (
	"testing"
	"github.com/asdfsx/codility/common"
)

func TestSolution4(t *testing.T) {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	result := Solution4(5, A)
	if !common.SliceCompare(result, []int{3, 2, 2, 4, 2}) {
		t.Errorf("expect [3, 2, 2, 4, 2], but found %s", result)
	}
}
