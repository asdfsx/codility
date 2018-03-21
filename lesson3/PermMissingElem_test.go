package lesson3

import (
	"testing"
)

func TestSolution3(t *testing.T) {
	A:= []int{2,3,1,5}
	result := Solution3(A)
	if result != 4{
		t.Errorf("expect 4, but found %s", result)
	}
}