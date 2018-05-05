package common

func SliceCompare(A []int, B []int) bool {
	if len(A) != len(B){
		return false
	}

	for i, v := range A{
		if v != B[i]{
			return false
		}
	}
	return true
}
