func sumGame(num string) bool {
	// Find where the right half starts.
	mid := len(num) / 2

	// Store sums of known digits in both halves.
	leftSum, rightSum := 0, 0

	// Count '?' characters in both halves.
	leftQuestion, rightQuestion := 0, 0

	// Scan every character once.
	for i := 0; i < len(num); i++ {
		if i < mid {
			// Update information for the left half.
			if num[i] == '?' {
				leftQuestion++
			} else {
				leftSum += int(num[i] - '0')
			}
		} else {
			// Update information for the right half.
			if num[i] == '?' {
				rightQuestion++
			} else {
				rightSum += int(num[i] - '0')
			}
		}
	}

	// If Bob cannot make this equality true, Alice wins.
	return 2*(leftSum-rightSum) != 9*(rightQuestion-leftQuestion)
}