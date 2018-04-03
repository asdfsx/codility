package lesson1

import (
	"fmt"
)

var Fortest = "fortest"

func Solution(N int) int{
	bytestring := fmt.Sprintf("%b", N);

	maxgap, currentgap := 0, 0;

	fmt.Println(bytestring);

	for _, c := range bytestring{
		if c == '1' && currentgap > 0{
			if currentgap > maxgap {
				maxgap = currentgap
			}
			currentgap = 0
		} else if c == '0' {
			currentgap++
		}
	}

	return maxgap
}