package lesson3

func Solution2(X int, Y int, D int) int{
	if (Y - X) % D == 0{
		return (Y - X) / D
	} else {
		return (Y - X) / D +1
	}
}