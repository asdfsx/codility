package lesson1

import (
	"testing"
)

func TestSolution(t *testing.T){
	result := Solution(1041)
	if result != 5 {
		t.Errorf("expect result 5 but %d found", result)
	}

	result = Solution(9)
	if result != 2 {
		t.Errorf("expect result 2 but %d found", result)
	}

	result = Solution(529)
	if result != 4 {
		t.Errorf("expect result 4 but %d found", result)
	}

	result = Solution(20)
	if result != 1 {
		t.Errorf("expect result 1 but %d found", result)
	}

	result = Solution(15)
	if result != 0 {
		t.Errorf("expect result 0 but %d found", result)
	}
}