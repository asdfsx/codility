package lesson7

import (
	"testing"
)

func TestSolution4(t *testing.T) {
	result := Solution4("(()(())())")
	if result != 1 {
		t.Errorf("expect 1, but found %s", result)
	}

	result = Solution4("())")
	if result != 0 {
		t.Errorf("expect 0, but found %s", result)
	}
}
