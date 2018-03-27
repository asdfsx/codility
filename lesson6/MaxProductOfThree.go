package lesson6

// https://codesays.com/2014/solution-to-max-product-of-three-by-codility/
// 如果全部是正数的话最大的3个的数乘积应该是最大的。如果最小的两个如果是负值的话，负负得正，可能可以算出一个最大值

func Solution3(A []int) int{
	max3 := [3]int{-2000,-2000,-2000}
	// Invariant: maxes[0] <= maxes[1] <= maxes[2]
	min2 := [2]int{2000,2000}
	// Invariant: mins[0] <= mins[1]

	for _, v := range A{
		if v >= max3[2]{
			max3[0] = max3[1]
			max3[1] = max3[2]
			max3[2] = v
		} else if v >= max3[1]{
			max3[0] = max3[1]
			max3[1] = v
		} else if v > max3[0]{
			max3[0] = v
		}
		if v <= min2[0] {
			min2[1] = min2[0]
			min2[0] = v
		} else if v < min2[1] {
			min2[1] = v
		}
	}
	if max3[0] * max3[1] * max3[2] > min2[0] * min2[1] * max3[2]{
		return max3[0] * max3[1] * max3[2]
	} else {
		return min2[0] * min2[1] * max3[2]
	}
	//sort.Ints(A)
	//length := len(A)
	//
	//last3 := A[length-3] * A[length-2] * A[length-1]
	//first2 := A[0] * A[1] * A[length-1]
	//
	//if last3 > first2{
	//	return last3
	//} else{
	//	return first2
	//}
}