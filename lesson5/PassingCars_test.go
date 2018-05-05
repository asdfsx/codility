package lesson5

import(
	"testing"
)

func TestSolution2(t *testing.T) {
	A := []int{0,1,0,1,1}
	result := Solution2(A)
	if result != 5{
		t.Errorf("expect 5, but found %s", result)
	}
}
