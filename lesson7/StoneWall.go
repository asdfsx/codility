package lesson7

// https://codesays.com/2014/solution-to-sigma-2012-stone-wall-by-codility/
// http://www.cnblogs.com/wei-li/p/StoneWall.html


func Solution(H []int) int {
	length := len(H)
	if length <= 1 {
		return length
	}

	stack := []int{}
	blockCount := 0

	for _, v := range H {
		for len(stack) > 0 && v < stack[len(stack)-1] {
			stack = stack[0:len(stack)-1]
			blockCount++
		}
		if len(stack) == 0 || v > stack[len(stack)-1] {
			stack = append(stack, v)
		}
	}

	blockCount += len(stack)
	return blockCount
}
