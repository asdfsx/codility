package lesson6

import (
	"sort"
	"fmt"
)

func Solution4(A []int) int{
	discsCount := len(A)                    // The total number of discs
	rangeUpper := make([]int, discsCount)   // The upper limit position of discs
	rangeLower := make([]int, discsCount)   // The lower limit position of discs

	// Fill the limit_upper and limit_lower
	for i := 0; i < discsCount; i++{
		rangeUpper[i] = i + A[i]
		rangeLower[i] = i - A[i]
	}

	sort.Ints(rangeUpper)
	sort.Ints(rangeLower)

	rangeLowerIndex := 0
	intersectCount := 0

	fmt.Println(rangeUpper, rangeLower)

	for rangeUpperIndex := 0; rangeUpperIndex < discsCount; rangeUpperIndex++{
		// Compute for each disc
		for rangeLowerIndex < discsCount && rangeUpper[rangeUpperIndex] >= rangeLower[rangeLowerIndex]{
				// Find all the discs that:
				//   disc_center - disc_radius <= current_center + current_radius
			rangeLowerIndex += 1
		}
		// We must exclude some discs such that:
		//    disc_center - disc_radius <= current_center + current_radius
		//    AND
		//    disc_center + disc_radius <(=) current_center + current_radius
		// These discs are not intersected with current disc, and below the
		//    current one completely.
		// After removing, the left discs are intersected with current disc,
		//    and above the current one.
		// Attention: the current disc is intersecting itself in this result
		//    set. But it should not be. So we need to minus one to fix it.
		intersectCount += rangeLowerIndex - rangeUpperIndex -1

		if intersectCount > 10000000{
			return -1
		}
	}

	return intersectCount
}