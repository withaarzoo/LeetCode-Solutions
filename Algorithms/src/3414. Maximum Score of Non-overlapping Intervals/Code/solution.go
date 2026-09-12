func maximumWeight(intervals [][]int) []int {
	type State struct {
		score  int64   // Total score of the selected intervals.
		ids    []int   // Original indices kept in sorted order.
		valid  bool    // True when this DP state is reachable.
	}

	n := len(intervals) // Number of intervals.
	const K = 4         // At most four intervals can be selected.

	// Store left, right, weight, and original index.
	a := make([][4]int64, n)

	// Save the original index before sorting.
	for i := 0; i < n; i++ {
		a[i] = [4]int64{
			int64(intervals[i][0]),
			int64(intervals[i][1]),
			int64(intervals[i][2]),
			int64(i),
		}
	}

	// Sort by right endpoint so binary search can find compatible prefixes.
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] < a[j][1]
	})

	// Store every right endpoint for predecessor searches.
	ends := make([]int64, n)
	for i := 0; i < n; i++ {
		ends[i] = a[i][1]
	}

	// dp[k][i] is the best result using exactly k intervals
	// from the first i sorted intervals.
	dp := make([][]State, K+1)
	for k := 0; k <= K; k++ {
		dp[k] = make([]State, n+1)
	}

	// Choosing zero intervals is always possible with score zero.
	for i := 0; i <= n; i++ {
		dp[0][i] = State{
			score: 0,
			ids:   []int{},
			valid: true,
		}
	}

	// Returns true when a is better than b.
	better := func(a State, b State) bool {
		// Any valid state beats an invalid state.
		if !a.valid {
			return false
		}
		if !b.valid {
			return true
		}

		// Higher score always wins.
		if a.score != b.score {
			return a.score > b.score
		}

		// Equal scores are compared lexicographically.
		for i := 0; i < len(a.ids) && i < len(b.ids); i++ {
			if a.ids[i] != b.ids[i] {
				return a.ids[i] < b.ids[i]
			}
		}

		// If one is a prefix of the other, shorter is smaller.
		return len(a.ids) < len(b.ids)
	}

	// Finds the first endpoint >= target.
	// Everything before that position has endpoint < target.
	lowerBound := func(length int, target int64) int {
		lo := 0
		hi := length

		// Standard lower-bound binary search.
		for lo < hi {
			mid := lo + (hi-lo)/2

			// End >= target is incompatible because touching overlaps.
			if ends[mid] >= target {
				hi = mid
			} else {
				// End < target is compatible, so move right.
				lo = mid + 1
			}
		}

		// This position is the number of compatible previous intervals.
		return lo
	}

	// Process intervals in sorted order.
	for i := 1; i <= n; i++ {
		left := a[i-1][0]         // Current interval's left endpoint.
		weight := a[i-1][2]       // Current interval's weight.
		originalIndex := int(a[i-1][3]) // Original index before sorting.

		// Find the compatible prefix before the current interval.
		p := lowerBound(i-1, left)

		// Try selecting exactly k intervals.
		for k := 1; k <= K; k++ {
			// Option 1: skip the current interval.
			dp[k][i] = dp[k][i-1]

			// Option 2: take the current interval.
			prev := dp[k-1][p]

			if prev.valid {
				// Copy the previous indices so other states are not modified.
				ids := append([]int{}, prev.ids...)

				// Add the current original index.
				ids = append(ids, originalIndex)

				// Sort the tiny index list for lexicographical comparison.
				sort.Ints(ids)

				// Build the candidate created by taking this interval.
				take := State{
					score: prev.score + weight,
					ids:   ids,
					valid: true,
				}

				// Keep the better of skip and take.
				if better(take, dp[k][i]) {
					dp[k][i] = take
				}
			}
		}
	}

	// Compare all possible final states containing at most four intervals.
	var answer State

	// Check every possible number of selected intervals.
	for k := 1; k <= K; k++ {
		if better(dp[k][n], answer) {
			answer = dp[k][n]
		}
	}

	// Return the original indices of the best choice.
	return answer.ids
}