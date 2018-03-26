package lesson5

import "fmt"

// https://codesays.com/2014/solution-to-min-avg-two-slice-by-codility/
// https://blog.csdn.net/xyqzki/article/details/50057717

func Solution4(A []int) int{
	min_avg_value := float32(A[0] + A[1])/2.0
	min_avg_pos := 0

	length := len(A)

	for i:=0;i< length-2;i++ {
		fmt.Println((A[i]+A[i+1])/2.0 )
		if float32(A[i]+A[i+1])/2.0 < min_avg_value {
			min_avg_value = float32(A[i] + A[i+1]) / 2.0
			min_avg_pos = i
		}
		if float32(A[i]+A[i+1]+A[i+2])/3.0 < min_avg_value {
			min_avg_value = float32(A[i] + A[i+1] + A[i+2]) / 3.0
			min_avg_pos = i
		}
	}


	if float32(A[length-2]+A[length-1])/2.0 < min_avg_value{
		min_avg_value = float32(A[length-2]+A[length-1])/2.0
		min_avg_pos = len(A)-2
	}

	return min_avg_pos
}