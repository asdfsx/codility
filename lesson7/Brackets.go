package lesson7

func Solution2(S string) int{
	if len(S) == 0{
		return 1
	}
	stack := []int32{}
	for _, v := range S{
		if len(stack) == 0 {
			stack = append(stack, v)
		} else if len(stack)>0 && v == ')' && stack[len(stack)-1] =='(' {
			stack = stack[0: len(stack)-1]
		} else if len(stack)>0 && v == ']' && stack[len(stack)-1] =='[' {

			stack = stack[0: len(stack)-1]
		} else if len(stack)>0 && v == '}' && stack[len(stack)-1] =='{' {
			stack = stack[0: len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	if len(stack) == 0{
		return 1
	} else {
		return 0
	}
}