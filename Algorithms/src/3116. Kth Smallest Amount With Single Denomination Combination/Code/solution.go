func findKthSmallest(coins []int, k int) int64 {
	// Sort coins so smaller denominations are processed first.
	sort.Ints(coins)

	// Keep only denominations that are not already covered.
	useful := make([]int64, 0)

	for _, coin := range coins {
		redundant := false

		// Check whether a smaller kept coin already covers all multiples
		// of the current coin.
		for _, prev := range useful {
			if int64(coin)%prev == 0 {
				redundant = true
				break
			}
		}

		// Keep only coins that add new possible amounts.
		if !redundant {
			useful = append(useful, int64(coin))
		}
	}

	m := len(useful)

	// The kth multiple of the smallest coin is always a valid upper bound.
	low := int64(1)
	high := useful[0] * int64(k)

	totalMasks := 1 << m

	// lcms stores the LCM of every non-empty subset.
	lcms := make([]int64, totalMasks)

	// signs stores +1 for odd subsets and -1 for even subsets.
	signs := make([]int64, totalMasks)

	// Precompute subset LCMs and inclusion-exclusion signs.
	for mask := 1; mask < totalMasks; mask++ {
		currentLCM := int64(1)
		bits := 0

		for i := 0; i < m; i++ {
			// Include useful[i] when its bit is set in this subset.
			if mask&(1<<i) != 0 {
				g := gcd(currentLCM, useful[i])

				// Divide before multiplying to calculate the LCM safely.
				currentLCM /= g

				// Cap the LCM if it becomes too large to contribute.
				if currentLCM > high/useful[i] {
					currentLCM = high + 1
					break
				}

				currentLCM *= useful[i]
				bits++
			}
		}

		// Save the LCM for this subset.
		lcms[mask] = currentLCM

		// Odd subsets are added and even subsets are subtracted.
		if bits%2 == 1 {
			signs[mask] = 1
		} else {
			signs[mask] = -1
		}
	}

	// Count how many valid amounts are less than or equal to x.
	count := func(x int64) int64 {
		result := int64(0)

		for mask := 1; mask < totalMasks; mask++ {
			// Ignore subsets whose LCM is larger than x.
			if lcms[mask] <= x {
				result += signs[mask] * (x / lcms[mask])
			}
		}

		return result
	}

	// Binary search for the kth smallest valid amount.
	for low < high {
		mid := low + (high-low)/2

		// Search left when mid already contains at least k valid amounts.
		if count(mid) >= int64(k) {
			high = mid
		} else {
			// Otherwise, search larger values.
			low = mid + 1
		}
	}

	// low is the kth smallest valid amount.
	return low
}

// gcd returns the greatest common divisor using the Euclidean algorithm.
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}