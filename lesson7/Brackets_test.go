package lesson7

import (
	"testing"
)

func TestSolution2(t *testing.T) {
	result := Solution2("{[()()]}")
	if result != 1 {
		t.Errorf("expect 1, but found %v", result)
	}

	result = Solution2("([)()]")
	if result != 0 {
		t.Errorf("expect 0, but found %v", result)
	}
}
