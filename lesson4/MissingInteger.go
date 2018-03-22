package lesson4

func Solution2(A []int) int{
	cache := map[int]int{}
	min := 0
	max := 0

	for i, v := range A{
		cache[v] = i
		if i == 0{
			min, max = v, v
		} else if min >= v {
			min = v
		} else if max <= v{
			max = v
		}
	}

	if min > 1 || max < 1{
		return 1
	}

	for i := 1; i < max; i++{
		if _, ok := cache[i]; !ok{
			return i
		}
	}
	return max + 1
}
