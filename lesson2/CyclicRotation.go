package lesson2

import "fmt"

func Solution(A []int, K int) []int{
	Alen := len(A)
	if Alen == 0 {
		return A
	}

	oldslice := A
	for i := 0; i < K; i++{
		newslice := append(oldslice[Alen-1:Alen], oldslice[0: Alen-1]...)
		oldslice = newslice
		fmt.Println(format(oldslice))
	}
	return oldslice
}

func format(A []int) (result string) {
    Alen := len(A)
    for i, v := range A{
    	if i == 0{
    		result = fmt.Sprintf("[%d, ", v)
		} else if i == Alen-1{
			result += fmt.Sprintf("%d]", v)
		} else{
			result  += fmt.Sprintf("%d, ", v)
		}
	}
	return
}