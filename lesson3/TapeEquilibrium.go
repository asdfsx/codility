package lesson3

import "fmt"

func Solution(A []int) int {
	cache := []int{}
	countall := 0;
	maxsize := 0;

	if len(A) == 1{
		return A[0]
	}

	for i, v := range A{
		maxsize = i
		countall += v
		cache = append(cache, countall)
		fmt.Println(countall, v)
	}

	result := 0;
	for i, v := range cache{
		if i == maxsize{
			break
		}
		tmp := 2 * v - countall
		if tmp >= 0 && (result > tmp || i == 0) {
			result = tmp
		} else if tmp < 0 && (result > -tmp || i == 0) {
			result = -tmp
		}
		fmt.Println(result)
	}

	return result
}