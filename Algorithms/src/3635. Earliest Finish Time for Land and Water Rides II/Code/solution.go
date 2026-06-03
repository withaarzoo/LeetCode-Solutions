func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {

	type Ride struct {
		start int
		dur   int
	}

	// Computes the best answer when category A is taken first
	var solve func([]int, []int, []int, []int) int64

	solve = func(startA []int, durA []int,
		startB []int, durB []int) int64 {

		m := len(startB)

		rides := make([]Ride, m)

		// Store (start, duration)
		for i := 0; i < m; i++ {
			rides[i] = Ride{startB[i], durB[i]}
		}

		// Sort by start time
		sort.Slice(rides, func(i, j int) bool {
			return rides[i].start < rides[j].start
		})

		starts := make([]int, m)
		prefixMinDur := make([]int64, m)
		suffixMinFinish := make([]int64, m)

		for i := 0; i < m; i++ {
			starts[i] = rides[i].start

			if i == 0 {
				prefixMinDur[i] = int64(rides[i].dur)
			} else {
				prefixMinDur[i] = min64(
					prefixMinDur[i-1],
					int64(rides[i].dur),
				)
			}
		}

		for i := m - 1; i >= 0; i-- {
			finish := int64(rides[i].start + rides[i].dur)

			if i == m-1 {
				suffixMinFinish[i] = finish
			} else {
				suffixMinFinish[i] = min64(
					suffixMinFinish[i+1],
					finish,
				)
			}
		}

		ans := int64(1 << 60)

		for i := 0; i < len(startA); i++ {

			// Finish time of first ride
			finish1 := int64(startA[i] + durA[i])

			// First index with start > finish1
			pos := sort.Search(m, func(j int) bool {
				return int64(starts[j]) > finish1
			})

			if pos > 0 {
				ans = min64(
					ans,
					finish1+prefixMinDur[pos-1],
				)
			}

			if pos < m {
				ans = min64(
					ans,
					suffixMinFinish[pos],
				)
			}
		}

		return ans
	}

	ans1 := solve(
		landStartTime,
		landDuration,
		waterStartTime,
		waterDuration,
	)

	ans2 := solve(
		waterStartTime,
		waterDuration,
		landStartTime,
		landDuration,
	)

	if ans1 < ans2 {
		return int(ans1)
	}
	return int(ans2)
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}