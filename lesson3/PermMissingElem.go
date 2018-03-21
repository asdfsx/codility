package lesson3

func Solution3(A []int) int{
	min := 1
	max := len(A)+1
	expectall := (min + max) * max / 2
	countall := 0

	for _, v:= range A{
		countall += v
	}

	return expectall - countall
}