package lesson6

import "sort"

// https://codesays.com/2014/solution-to-min-avg-two-slice-by-codility/
// 数组里任意两个元素相加大于第三个
// 在排序后的数组里问题被简化：a[1],a[2],a[3]三个元素 a[3]与任意一个相加都肯定是最大的，只需要找出满足a[1]+a[2]>a[3]的就可以了

func Solution(A []int) int{
	if len(A) < 3{
		return 0
	}

	sort.Ints(A)
	length := len(A)

	for i:= 0; i<length-2; i++{
		if A[i]+A[i+1] > A[i+2]{
			return 1
		}
	}

	return 0
}