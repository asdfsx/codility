package lesson7

func Solution3(A []int, B []int) int{
	aliveCount := 0         // The number of fish that will stay alive
	downstream := []int{}   // To record the fishs flowing downstream
	downstreamCount := 0    // To record the number of elements in downstream

	for i, v := range A{
		if B[i] == 1{
			// This fish is flowing downstream. It would
			// NEVER meet the previous fishs. But possibly
			// it has to fight with the downstream fishs.
			downstream = append(downstream, v)
			downstreamCount += 1
		} else {
			// This fish is flowing upstream. It would either
			//    eat ALL the previous downstream-flow fishs,
			//    and stay alive.
			// OR
			//    be eaten by ONE of the previous downstream-
			//    flow fishs, which is bigger, and died.
			for downstreamCount != 0 {
				// It has to fight with each previous living
				// fish, with nearest first.
				if downstream[len(downstream)-1] < A[i] {
					// Win and to continue the next fight
					downstreamCount -= 1
					downstream = downstream[0:len(downstream)-1]
				} else {
					// Lose and die
					break
				}
			}
			// This upstream-flow fish eat all the previous
			// downstream-flow fishs. Win and stay alive.
			if downstreamCount == 0{
			    aliveCount += 1
			}

		}
	}

	// Currently, all the downstream-flow fishs in stack
	// downstream will not meet with any fish. They will
	// stay alive.
	aliveCount += len(downstream)
	return aliveCount
}
