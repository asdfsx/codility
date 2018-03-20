package lesson2

func Solution2(A []int) int{

	result := 0

	for _, v := range A {
		result ^= v
	}

	return result;
}