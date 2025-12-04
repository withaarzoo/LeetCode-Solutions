func countCollisions(directions string) int {
	n := len(directions)
	i, j := 0, n-1

	// Skip leading 'L' cars (no collision)
	for i < n && directions[i] == 'L' {
		i++
	}

	// Skip trailing 'R' cars (no collision)
	for j >= 0 && directions[j] == 'R' {
		j--
	}

	collisions := 0
	// Every non-'S' in this middle segment will collide exactly once
	for k := i; k <= j; k++ {
		if directions[k] != 'S' {
			collisions++
		}
	}

	return collisions
}
