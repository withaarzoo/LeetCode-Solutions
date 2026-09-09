func countCommas(n int64) int64 {
	// Start at 1000 because numbers smaller than 1000 have no commas.
	start := int64(1000)

	// The first comma group contains one comma per number.
	commas := int64(1)

	// Store the total number of commas.
	var answer int64

	// Process every comma group that starts within the range.
	for start <= n {
		var end int64

		// If the next group boundary is beyond n, the current group ends at n.
		if start > n/1000 {
			end = n
		} else {
			// Otherwise, the current group ends immediately before start * 1000.
			end = start*1000 - 1
		}

		// Count how many numbers are in the current group.
		count := end - start + 1

		// Every number in the group has the same comma count.
		answer += count * commas

		// Move to the next group of numbers.
		start *= 1000

		// The next group has one additional comma.
		commas++
	}

	// Return the total number of commas.
	return answer
}