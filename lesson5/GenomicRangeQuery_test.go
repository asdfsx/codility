package lesson5

import(
	"testing"
	"github.com/asdfsx/codility/common"
)

func TestSolution3(t *testing.T) {
	S := ""
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}
	result := Solution3(S, P, Q)
	if !common.SliceCompare(result, []int{2, 4, 1}){
		t.Errorf("expect [2, 4, 1], but found %v", result)
	}
}
