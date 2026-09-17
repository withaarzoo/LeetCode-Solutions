func minSumOfLengths(arr []int, target int) int {
	n := len(arr)

	// best[i] stores the shortest target-sum sub-array
	// found completely inside arr[0..i].
	const INF = int(^uint(0) >> 1)
	best := make([]int, n)

	// Fill best with INF to represent "no valid sub-array yet".
	for i := 0; i < n; i++ {
		best[i] = INF
	}

	left := 0      // Left boundary of the sliding window.
	sum := 0       // Sum of the current sliding window.
	answer := INF  // Minimum combined length found so far.

	for right := 0; right < n; right++ {
		// Expand the window by adding arr[right].
		sum += arr[right]

		// All elements are positive, so shrinking from the left
		// always makes the window sum smaller.
		for sum > target {
			sum -= arr[left]
			left++
		}

		// The current window has exactly the target sum.
		if sum == target {
			currentLength := right - left + 1

			// best[left-1] represents a valid sub-array
			// that ends before the current window starts.
			if left > 0 && best[left-1] != INF {
				candidate := currentLength + best[left-1]

				// Keep the smallest combined length.
				if candidate < answer {
					answer = candidate
				}
			}

			// Store the current valid sub-array length.
			best[right] = currentLength
		}

		// Carry the best answer from the previous prefix.
		if right > 0 && best[right-1] < best[right] {
			best[right] = best[right-1]
		}
	}

	// Return -1 when two non-overlapping target-sum sub-arrays do not exist.
	if answer == INF {
		return -1
	}

	return answer
}