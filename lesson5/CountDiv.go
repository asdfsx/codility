package lesson5

// 计算的时候别使用因式合并，因为计算机里有精度问题

func Solution(A int, B int, K int) int {
	//if A == B && A % K == 0{
	//	return 1
	//}
	//if K >= A && A != 0 {
	//	A = 1
	//}
	//fmt.Println(B, A, K, (B - A) / K, B/K - A/K)
	//if r := A % K; r == 0{
	//	return (B - A) / K + 1
	//}
	//return (B - A) / K
	if r := A % K; r == 0 {
		return B/K - A/K + 1
	} else {
		return B/K - A/K + 0
	}
}
