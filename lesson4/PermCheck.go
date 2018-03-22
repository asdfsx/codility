package lesson4

func Solution3(A []int) int{
	var cache []int = make([]int, len(A))
	limit := len(A)
    for _, v := range A {
		if v < 1 && v> limit{
			return 0
		}else if cache[v] == 1{
    		return 0
		} else {
			cache[v] = 1
		}
	}
	return 1
}

