func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// Store the reserved seats of each affected row as a bitmask.
	rows := make(map[int]int)

	// Process every reserved seat once.
	for _, seat := range reservedSeats {
		row := seat[0]
		col := seat[1]

		// Seats 1 and 10 cannot belong to any valid four-seat block.
		if col >= 2 && col <= 9 {
			// Set the bit corresponding to the reserved seat.
			rows[row] |= 1 << col
		}
	}

	// Every row not stored in the map has no relevant reservations.
	// Each such row can always fit two groups.
	answer := 2 * (n - len(rows))

	// Create masks for the three possible group blocks.
	// left = seats 2-5, middle = seats 4-7, right = seats 6-9.
	left := (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)
	middle := (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)
	right := (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)

	// Check only rows that contain relevant reserved seats.
	for _, mask := range rows {
		// A block is free when none of its seats are reserved.
		canLeft := (mask & left) == 0
		canMiddle := (mask & middle) == 0
		canRight := (mask & right) == 0

		// Left and right blocks do not overlap, so both groups can fit.
		if canLeft && canRight {
			answer += 2
		} else if canLeft || canMiddle || canRight {
			// If any block is free, I can place one group.
			answer++
		}
		// If all three blocks are blocked, this row gets zero groups.
	}

	// Return the maximum number of groups.
	return answer
}