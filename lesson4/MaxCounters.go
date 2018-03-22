package lesson4

func Solution4(N int, A []int) []int{
	counter := make([]int, N)
	currentmax := 0

	for _, X := range A {
		if X >= 1 && X <= N{
			currentmax = increase(X, counter, currentmax)
		} else if X == 1 + N {
			maxcounter(currentmax, counter)
		}
	}

	return counter
}

func increase(X int, counter []int, currentmax int) int {
	counter[X-1]++
	if counter[X-1] > currentmax{
		return counter[X-1]
	} else{
		return currentmax
	}
}

func maxcounter(max int, counter[]int){
    for i, _ := range counter{
    	counter[i] = max
	}
}