package lesson6

import "sort"

func Solution2(A []int) int{
	sort.Ints(A)
	currentValue := 0
	counter := 0

	for i, v := range A{
		if i == 0 && v == 0{
			counter ++
		} else if currentValue != v{
			counter ++
			currentValue = v
		}
	}

	return counter
}