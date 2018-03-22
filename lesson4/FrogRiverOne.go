package lesson4

func Solution(X int, A []int) int {
	leaves := map[int]int{}
	count := 0
	for i, v := range A{
		if _, ok := leaves[v]; !ok {
			leaves[v] = i
			count += 1
		}
		if count == X {
			return i
		}
	}

	return -1
}