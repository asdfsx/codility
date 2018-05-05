package lesson5

func Solution2(A []int) int{
	countzero, count := 0, 0

	for _, v := range A{
		if v == 0{
			countzero ++
		}else if v == 1{
			count += countzero
			if count > 1000000000{
				return -1
			}
		}
	}
	return count
}