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

// copy from https://www.martinkysel.com/codility-maxcounters-solution/
func Solution5(N int, A[]int) []int {
	sol := make([]int, N)
	current_max := 0
	last_increase := 0

	for i, v := range A{
		if v > N{
			last_increase = current_max;
		} else {
			sol[A[i]-1] = max(sol[A[i]-1], last_increase);
			sol[A[i]-1]++;
			current_max = max(current_max, sol[A[i]-1]);
		}
	}

	for i := 0; i < N; i++ {
	    sol[i] = max(sol[i], last_increase);
	}
	return sol
}

func max(a int, b int) int{
	if a > b {
		return a
	} else {
		return b
	}
}