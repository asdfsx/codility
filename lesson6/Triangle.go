package lesson6

import "sort"

func Solution(A []int) int{
	if len(A) < 3{
		return 0
	}

	sort.Ints(A)
	length := len(A)

	for i:= 0; i<length-2; i++{
		if A[i]+A[i+1] > A[i+2]{
			return 1
		}
	}

	return 0
}