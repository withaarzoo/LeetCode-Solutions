func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)

	// Store robot indices
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	// Sort indices based on positions
	sort.Slice(indices, func(i, j int) bool {
		return positions[indices[i]] < positions[indices[j]]
	})

	// Stack to store robots moving right
	stack := []int{}

	for _, idx := range indices {
		// Robot moving right
		if directions[idx] == 'R' {
			stack = append(stack, idx)
		} else {
			// Robot moving left
			for len(stack) > 0 && healths[idx] > 0 {
				topIdx := stack[len(stack)-1]

				if healths[topIdx] < healths[idx] {
					stack = stack[:len(stack)-1]
					healths[idx]--
					healths[topIdx] = 0
				} else if healths[topIdx] == healths[idx] {
					stack = stack[:len(stack)-1]
					healths[topIdx] = 0
					healths[idx] = 0
				} else {
					healths[topIdx]--
					healths[idx] = 0
				}
			}
		}
	}

	// Collect surviving robots
	result := []int{}
	for _, health := range healths {
		if health > 0 {
			result = append(result, health)
		}
	}

	return result
}